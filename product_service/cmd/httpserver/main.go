package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/meokg456/productservice/adapter/dynamostore"
	"github.com/meokg456/productservice/adapter/grpcserver"
	"github.com/meokg456/productservice/adapter/httpserver"
	"github.com/meokg456/productservice/adapter/kafkaservice"
	"github.com/meokg456/productservice/pkg/config"
	"github.com/meokg456/productservice/pkg/logger"
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

	options := dynamostore.ParseFromConfig(config)
	db, err := dynamostore.NewConnection(options)
	if err != nil {
		applog.Fatalf("Cannot connect to db %v", err)
	}

	productBrokerOptions := kafkaservice.ParseFromConfig(config)
	kafkaWriter := kafkaservice.NewWriter(productBrokerOptions)
	productBroker := kafkaservice.NewProductBroker(kafkaWriter)

	server := httpserver.New(config)

	productStore := dynamostore.NewProductStore(db)

	server.Logger = applog
	server.ProductStore = productStore

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		applog.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpcserver.New(config)
	grpcServer.Logger = applog
	grpcServer.ProductStore = productStore
	grpcServer.ProductBroker = productBroker

	go func() {
		applog.Info("Grpc server started!")
		applog.Fatal(grpcServer.Serve(listener))
	}()

	applog.Info("Server started!")
	applog.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), server))
}
