package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/meokg456/inventoryservice/adapter/dynamostore"
	"github.com/meokg456/inventoryservice/adapter/grpcserver"
	"github.com/meokg456/inventoryservice/adapter/httpserver"
	"github.com/meokg456/inventoryservice/pkg/config"
	"github.com/meokg456/inventoryservice/pkg/logger"
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

	server := httpserver.New(config)

	inventoryStore := dynamostore.NewInventoryStore(db)

	server.Logger = applog

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		applog.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpcserver.New(config)
	grpcServer.Logger = applog
	grpcServer.InventoryStore = inventoryStore

	go func() {
		applog.Fatal(grpcServer.Serve(listener))
	}()

	applog.Info("Server started!")
	applog.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), server))
}
