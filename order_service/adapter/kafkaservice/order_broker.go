package kafkaservice

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
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

func (o *OrderBroker) SaveOrder(ord *order.Order) error {
	if ord.Id == "" {
		ord.Id = uuid.NewString()
		ord.Status = order.Pending
	}

	items := []ItemMessage{}
	for _, item := range ord.Items {
		items = append(items, ItemMessage{
			ProductId: item.ProductId,
			Types:     item.Types,
			Quantity:  item.Quantity,
		})
	}

	orderMessage := OrderMessage{
		Id:      ord.Id,
		UserId:  ord.UserId,
		Status:  string(ord.Status),
		Payment: string(ord.Payment),
		Paid:    ord.Paid,
		Items:   items,
	}

	message, err := json.Marshal(orderMessage)
	if err != nil {
		return err
	}

	err = o.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(ord.Id),
		Value: message,
	})

	if err != nil {
		return err
	}

	return nil
}
