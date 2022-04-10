package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

var Address3 = []string{""}

func main() {
	syncProducer3(Address3)
}

//同步消息模式
func syncProducer3(address []string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer func(p sarama.SyncProducer) {
		err := p.Close()
		if err != nil {
		}
	}(p)
	topic := ""
	writeToKafka := ``
	if err != nil{
		fmt.Println(err)
		return
	}
	//var event proto.Message
	//err = json.Unmarshal([]byte(writeToKafka), event)
	//a,err := proto.Marshal(event)
	//fmt.Println(err)
	//msg := &sarama.ProducerMessage{
	//	Topic: topic,
	//	Value: sarama.ByteEncoder(a),
	//}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(writeToKafka),
	}
	fmt.Print(msg)
	part, offset, err := p.SendMessage(msg)
	if err != nil {
		log.Printf("send message(%s) err=%s \n", msg, err)
	} else {
		fmt.Fprintf(os.Stdout, "发送成功，partition=%d, offset=%d \n", part, offset)
	}
}
