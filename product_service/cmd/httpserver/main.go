package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	pb "proto/product"

	"github.com/meokg456/productservice/adapter/dynamostore"
	"github.com/meokg456/productservice/adapter/grpcserver"
	"github.com/meokg456/productservice/adapter/httpserver"
	"github.com/meokg456/productservice/pkg/config"
	"github.com/meokg456/productservice/pkg/logger"
	"google.golang.org/grpc"
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

	productStore := dynamostore.NewProductStore(db)

	server.Logger = applog
	server.ProductStore = productStore

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		applog.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpcService := grpcserver.New(config)
	grpcService.Logger = applog
	grpcService.ProductStore = productStore

	pb.RegisterProductServiceServer(grpcServer, grpcService)
	go func() {
		applog.Fatal(grpcServer.Serve(listener))
	}()

	applog.Info("Server started!")
	applog.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), server))
}
