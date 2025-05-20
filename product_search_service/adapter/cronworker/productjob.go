package cronworker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/meokg456/productsearchservice/adapter/cronworker/model"
	"github.com/meokg456/productsearchservice/domain/product"
)

func (c *CronWorker) AddProductJob() {
	c.Cron.AddFunc("@daily", c.UpdateProduct)
}

func (c *CronWorker) UpdateProduct() {

	now := time.Now()
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		m, err := c.ProductReader.FetchMessage(ctx)
		if err != nil {
			c.Logger.Errorf("cron update product: failed to read messages product broker %v", err)
			break
		}

		if m.Time.After(now) {
			break
		}

		var pro model.ProductMessage
		err = json.Unmarshal(m.Value, &pro)
		if err != nil {
			c.ProductReader.CommitMessages(context.Background(), m)
			c.Logger.Errorf("cron update product: failed to unmarshal product value of %v, error: %v", m.Value, err)
			continue
		}

		err = c.ProductStore.SaveProduct(product.Product{
			Id:           pro.Id,
			Title:        pro.Title,
			Descriptions: pro.Descriptions,
			Category:     pro.Category,
			Images:       pro.Images,
			AdditionInfo: pro.AdditionInfo,
			MerchantId:   pro.MerchantId,
		})
		if err != nil {
			c.Logger.Errorf("cron update product: failed to save product %v, error: %v", pro, err)
			continue
		}
		err = c.ProductReader.CommitMessages(context.Background(), m)
		if err != nil {
			c.Logger.Errorf("cron update product: failed to commit message %v", err)
		}
	}
}
