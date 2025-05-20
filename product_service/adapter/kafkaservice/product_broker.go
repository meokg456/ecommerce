package kafkaservice

import (
	"context"
	"encoding/json"

	"github.com/meokg456/productservice/domain/product"
	"github.com/segmentio/kafka-go"
)

type ProductBroker struct {
	writer *kafka.Writer
}

func NewProductBroker(writer *kafka.Writer) *ProductBroker {
	return &ProductBroker{
		writer: writer,
	}
}

func (p *ProductBroker) SendProductChange(pro product.Product) error {
	data, err := json.Marshal(ProductMessage{
		Id:           pro.Id,
		Title:        pro.Title,
		Descriptions: pro.Descriptions,
		Category:     pro.Category,
		Images:       pro.Images,
		AdditionInfo: pro.AdditionInfo,
		MerchantId:   pro.MerchantId,
	})

	if err != nil {
		return err
	}

	err = p.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(pro.Id),
		Value: data,
	})

	if err != nil {
		return err
	}

	return nil
}
