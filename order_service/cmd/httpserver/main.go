package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/meokg456/orderservice/adapter/httpserver"
	"github.com/meokg456/orderservice/adapter/kafkaservice"
	"github.com/meokg456/orderservice/pkg/config"
	"github.com/meokg456/orderservice/pkg/logger"
)

func main() {
	applog, err := logger.NewAppLogger()
	if err != nil {
		log.Fatalf("Cannot load config %v", err)
	}
	defer applog.Sync()

	config, err := config.LoadConfig()
	if err != nil {
		applog.Fatalf("Cannot load env config %v", err)
	}

	productBrokerOptions := kafkaservice.ParseFromConfig(config)
	kafkaWriter := kafkaservice.NewWriter(productBrokerOptions)
	orderBroker := kafkaservice.NewOrderBroker(kafkaWriter)

	server := httpserver.New(config)

	server.Logger = applog
	server.OrderBroker = orderBroker

	applog.Info("Server started!")
	applog.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), server))
}
