package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

const (
	MeetingRoomStatusFree   = "free"
	MeetingRoomStatusBusy   = "busy"
	MeetingRoomStatusClosed = "closed"

	UserToken = "11e1fde8f60b73323e4299057a5dd66b5be74199"
)

type MeetingRoom struct {
	sync.RWMutex
	ID       string
	Status   string
	Definite bool
	LastUpdatedTime time.Time
	LastBookByMeTime time.Time

	cancelInfo context.CancelFunc
	cancelBook context.CancelFunc
}

var meetingRooms = make(map[string]*MeetingRoom)
var client = &http.Client{
	Timeout: time.Second * 60,
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 300 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          150,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   100,
	},
}
var nullTime = time.Time{}

func init() {
	for i := 1; i <= 30; i++ {
		meetingRooms[strconv.Itoa(i)] = &MeetingRoom{
			ID:       strconv.Itoa(i),
			Status:   MeetingRoomStatusFree,
			Definite: true,
		}
	}
}

func main() {
	kafkaChan := make(chan string)
	go syncKafkaMessage(kafkaChan)
	go Consume(kafkaChan)
	go roundRobinIndefinite()
	go roundRobinBook()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
}

func syncKafkaMessage(roomIdChan chan string) {
	for id := range roomIdChan {
		room := meetingRooms[id]
		now := time.Now()
		room.Lock()
		sub := now.Sub(room.LastBookByMeTime)
		if sub > time.Millisecond * 300 || sub < time.Millisecond * -300 {
			room.LastBookByMeTime = nullTime
			room.Definite = false
		}
		room.LastUpdatedTime = now
		room.Unlock()
	}
}

func roundRobinIndefinite() {
	wp := NewWorkerPool()
	for {
		time.Sleep(50 * time.Millisecond)
		for _, r := range meetingRooms {
			room := r
			room.RLock()
			if !room.Definite && room.cancelInfo == nil {
				wp.Submit(func() {
					room.Lock()
					if room.cancelInfo != nil {
						room.Unlock()
						return
					}
					ctx, cancel := context.WithCancel(context.Background())
					room.cancelInfo = cancel
					room.Unlock()
					status, err := info(ctx, room.ID)
					room.Lock()
					if err == nil {
						room.Status = status
						room.Definite = true
						if status == MeetingRoomStatusClosed && room.cancelBook != nil {
							room.cancelBook()
						}
					}
					room.cancelInfo = nil
					room.Unlock()
				})
			}
			room.RUnlock()
		}
	}
}

func roundRobinBook() {
	wp := NewWorkerPool()
	for {
		time.Sleep(50*time.Millisecond)
		var bestRoomId string
		var maxScore int
		for _, room := range meetingRooms {
			score := func() int {
				room.RLock()
				defer room.RUnlock()
				if !room.Definite {
					if room.Status == MeetingRoomStatusBusy {
						return 50
					}
					return 30
				}
				now := time.Now()
				if room.Status == MeetingRoomStatusBusy {
					sub := now.Sub(room.LastUpdatedTime)
					if room.LastBookByMeTime != nullTime {
						if sub > 10 * time.Second {
							return 100
						} else if sub > 5 * time.Second {
							return 70
						} else {
							return 30
						}
					} else {
						if sub > 10 * time.Second {
							return 80
						} else if sub > 5 * time.Second {
							return 50
						} else {
							return 30
						}
					}
				}
				if room.Status == MeetingRoomStatusFree {
					if room.LastBookByMeTime != nullTime {
						return 70
					} else {
						return 50
					}
				}
				if room.Status == MeetingRoomStatusClosed {
					return 30
				}
				return 30
			} ()
			if score > maxScore {
				maxScore = score
				bestRoomId = room.ID
			}
		}
		wp.Submit(func() {
			room := meetingRooms[bestRoomId]
			room.Lock()
			if room.cancelBook != nil {
				room.Unlock()
				return
			}
			ctx, cancel := context.WithCancel(context.Background())
			room.cancelBook = cancel
			room.Unlock()
			success, err := book(ctx, room.ID)
			room.Lock()
			if err == nil {
				if success {
					fmt.Println("success order", room.ID)
					room.LastBookByMeTime = time.Now()
				} else {
					fmt.Println("failed order", room.ID)
				}
				room.Definite = true
				room.Status = MeetingRoomStatusBusy
				if room.cancelInfo != nil {
					room.cancelInfo()
				}
			}
			room.cancelBook = nil
			room.Unlock()
		})
	}
}

func info(ctx context.Context, id string) (string, error) {
	body := map[string]string{
		"meeting_room_id": id,
		"user_token": UserToken,
	}
	reqBody, _ := json.Marshal(body)
	request, err := http.NewRequestWithContext(ctx, "POST", "http://10.30.102.148:80/api/info", bytes.NewReader(reqBody))
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return string(respBody), nil
	}
	return "", errors.New("error code")
}

func book(ctx context.Context, id string) (bool, error) {
	body := map[string]string{
		"meeting_room_id": id,
		"user_token": UserToken,
	}
	reqBody, _ := json.Marshal(body)
	request, err := http.NewRequestWithContext(ctx, "POST", "http://10.30.102.148:80/api/book", bytes.NewReader(reqBody))
	resp, err := client.Do(request)
	if err != nil {
		return false, err
	}
	return resp.StatusCode == 200, nil
}

// ------------------------------------------ kafka ------------------------------------------

type RoomMessage struct {
	MeetingRoomId int `json:"meeting_room_id"`
}

type consumerGroupHandler struct {
	msgChan chan string
}

func (*consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (*consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		message := &RoomMessage{}
		_ = json.Unmarshal(msg.Value, message)
		h.msgChan <- strconv.Itoa(message.MeetingRoomId)
		// 手动确认消息
		sess.MarkMessage(msg, "")
	}
	return nil
}

func Consume(msgChan chan string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V1_0_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumer, err := sarama.NewConsumerGroup([]string{"10.30.102.60:9092"}, "momo-q1c25", config)
	if err != nil {
		panic(err)
	}
	go func() {
		defer consumer.Close()
		for {
			err := consumer.Consume(context.Background(), []string{"meeting_status_channel"}, &consumerGroupHandler{msgChan: msgChan})
			if err != nil {
				panic(err)
			}
		}
	}()
}

// ------------------------------------------ workerpool ------------------------------------------

type WorkerPool struct {
	Concurrency int
	taskChan    chan func()
}

func NewWorkerPool() *WorkerPool {
	wp := &WorkerPool{
		Concurrency: 3,
		taskChan:    make(chan func()),
	}
	wp.Start()
	return wp
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.Concurrency; i++ {
		go func() {
			for task := range wp.taskChan {
				task()
			}
		}()
	}
}

func (wp *WorkerPool) Submit(task func()) {
	select {
	case wp.taskChan <- task:
	default:
	}
}

