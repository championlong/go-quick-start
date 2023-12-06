package kafka_consumer

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/championlong/go-quick-start/pkg/recovery"
	"github.com/wvanbergen/kafka/consumergroup"
	"golang.org/x/sync/semaphore"
)

const (
	BOOKING_STATUS_CG    = ""
	BOOKING_STATUS_TOPIC = "meeting_status_channel"
)

var zks = []string{
	"127.0.0.1:2181",
}

type roomBookingKafkaConsumer struct {
	consumerGroup *consumergroup.ConsumerGroup
	sema          *semaphore.Weighted
}

func (consumer *roomBookingKafkaConsumer) init() error {
	consumer.sema = semaphore.NewWeighted(3)
	cfg := consumergroup.NewConfig()
	cfg.Offsets.Initial = sarama.OffsetNewest
	cfg.Offsets.ProcessingTimeout = 10 * time.Second

	var err error
	consumer.consumerGroup, err = consumergroup.JoinConsumerGroup(
		BOOKING_STATUS_CG,
		[]string{BOOKING_STATUS_TOPIC},
		zks,
		cfg,
	)
	if err != nil {
		return err
	} else {
		fmt.Println("%s: join consumer group %s successfully", BOOKING_STATUS_TOPIC, BOOKING_STATUS_CG)
	}
	go func() {
		for err := range consumer.consumerGroup.Errors() {
			fmt.Println("consumer group error: %s", err.Error())
		}
	}()
	return nil
}

// for循环消费.
func (consumer *roomBookingKafkaConsumer) consume(ctx context.Context) {
ConsumeMessage:
	for {
		select {
		case <-ctx.Done():
			if err := consumer.consumerGroup.Close(); err != nil {
				fmt.Println("marketing_user_activation: error closing the consumer: %s", err.Error())
			}
			break ConsumeMessage
		case msg := <-consumer.consumerGroup.Messages():

			go consumer.ProcessMsg(ctx, msg.Value)

			// commit after process, confirm at least once
			err := consumer.consumerGroup.CommitUpto(msg)
			if err != nil {
				fmt.Println("%s: error committing: %s", BOOKING_STATUS_TOPIC, err.Error())
			}
		}
	}
}

func (c *roomBookingKafkaConsumer) ProcessMsg(ctx context.Context, msg []byte) {
	topic := BOOKING_STATUS_TOPIC
	fmt.Println("%s: %s", topic, msg)
}

func RunBookingStatusConsumer() {
	defer recovery.Recovery("run consumer")

	kafkaConsumer := &roomBookingKafkaConsumer{}

	ctx := context.Background()
	err := kafkaConsumer.init()
	if err != nil {
		fmt.Printf("consumer init error: %s", err.Error())
		return
	}
	kafkaConsumer.consume(ctx)
}
