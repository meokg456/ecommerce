package kafkaservice

import (
	"github.com/meokg456/orderservice/domain/order"
	"github.com/segmentio/kafka-go"
)

type OrderBroker struct {
	writer *kafka.Writer
}

func NewOrderBroker(writer *kafka.Writer) *OrderBroker {
	return &OrderBroker{
		writer: writer,
	}
}

func (o *OrderBroker) SaveOrder(order order.Order) error {
	return nil
}
