package cronworker

import (
	"github.com/meokg456/productsearchservice/domain/product"
	"github.com/meokg456/productsearchservice/pkg/config"
	"github.com/robfig/cron"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type CronWorker struct {
	Cron   *cron.Cron
	Logger *zap.SugaredLogger
	Config *config.Config

	ProductReader *kafka.Reader
	ProductConn   *kafka.Conn

	ProductStore product.Storage
}

func New(cfg *config.Config) *CronWorker {
	c := cron.New()

	productReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{cfg.MessageBroker.ProductBrokerHost},
		Topic:    cfg.MessageBroker.ProductTopic,
		GroupID:  cfg.MessageBroker.ProductGroupId,
		MaxBytes: 10e6,
	})

	worker := &CronWorker{
		Cron:          c,
		Config:        cfg,
		ProductReader: productReader,
	}

	worker.AddProductJob()

	return worker
}

func (c *CronWorker) Start() {
	c.Cron.Start()
}
