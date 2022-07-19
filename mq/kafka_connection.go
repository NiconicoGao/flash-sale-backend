package mq

import (
	"flash-sale-backend/utils"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

func KafkaInit(config *utils.KafkaConfig) error {
	var err error
	addr := fmt.Sprintf("%v:%v", config.IP, config.Port)
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": addr})
	if err != nil {
		return err
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": addr,
		"group.id":          config.GroupName,
	})
	if err != nil {
		return err
	}

	err = consumer.SubscribeTopics([]string{config.Topic}, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			msg, err := consumer.ReadMessage(-1)
			if err == nil {
				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			} else {
				// The client will automatically try to recover from all errors.
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}()
	return nil
}
