package main

import (
	"log"

	"github.com/meokg456/orderservice/pkg/config"
	"github.com/segmentio/kafka-go"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load env configs %v", err)
	}

	conn, err := kafka.Dial("tcp", config.MessageBroker.OrderBrokerHost)
	if err != nil {
		log.Fatalf("Cannot connect to message broker %v", err)
	}

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             config.MessageBroker.OrderTopic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		log.Fatalf("Cannot create topic %v", err)
	}
}
