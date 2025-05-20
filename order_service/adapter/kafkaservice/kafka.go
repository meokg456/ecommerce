package kafkaservice

import (
	"github.com/meokg456/orderservice/pkg/config"
	"github.com/segmentio/kafka-go"
)

type Options struct {
	Hosts []string
	Topic string
}

func ParseFromConfig(cfg *config.Config) Options {
	return Options{
		Hosts: []string{cfg.MessageBroker.OrderBrokerHost},
		Topic: cfg.MessageBroker.OrderTopic,
	}
}

func NewWriter(options Options) *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP(options.Hosts...),
		Topic: options.Topic,
	}
}
