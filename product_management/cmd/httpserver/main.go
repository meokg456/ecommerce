package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/meokg456/productmanagement/adapter/grpcservice"
	"github.com/meokg456/productmanagement/adapter/httpserver"
	"github.com/meokg456/productmanagement/adapter/postgresstore"
	"github.com/meokg456/productmanagement/pkg/config"
	"github.com/meokg456/productmanagement/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	inventorypb "github.com/meokg456/ecommerce/proto/inventory"
	productpb "github.com/meokg456/ecommerce/proto/product"
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

	postgresOptions := postgresstore.ParseFromConfig(config)
	db, err := postgresstore.NewConnection(postgresOptions)
	if err != nil {
		applog.Fatalf("Cannot connect to db %v", err)
	}

	productConn, err := grpc.NewClient(config.GrpcService.ProductGrpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		applog.Fatalf("Cannot create product grpc client", err)
	}

	inventoryConn, err := grpc.NewClient(config.GrpcService.InventoryGrpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		applog.Fatalf("Cannot create grpc client", err)
	}

	productClient := productpb.NewProductServiceClient(productConn)
	inventoryClient := inventorypb.NewInventoryServiceClient(inventoryConn)

	server := httpserver.New(config)

	userStore := postgresstore.NewUserStore(db)

	server.Logger = applog
	server.UserStore = userStore
	server.ProductService = grpcservice.NewProductService(productClient)
	server.InventoryService = grpcservice.NewInventoryService(inventoryClient)

	applog.Info("Server started!")
	applog.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), server))
}
