package kafka_consumer

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/Shopify/sarama"
)

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
		panic(recover())
	}
	go func() {
		defer consumer.Close()
		for {
			err := consumer.Consume(
				context.Background(),
				[]string{"meeting_status_channel"},
				&consumerGroupHandler{msgChan: msgChan},
			)
			if err != nil {
				panic(recover())
			}
		}
	}()
}
