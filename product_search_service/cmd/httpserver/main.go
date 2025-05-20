package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/meokg456/productsearchservice/adapter/cronworker"
	"github.com/meokg456/productsearchservice/adapter/httpserver"
	"github.com/meokg456/productsearchservice/adapter/postgresstore"
	"github.com/meokg456/productsearchservice/pkg/config"
	"github.com/meokg456/productsearchservice/pkg/logger"
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

	server := httpserver.New(config)

	productStore := postgresstore.NewProductStore(db)

	server.Logger = applog
	server.ProductStore = productStore

	cronWorker := cronworker.New(config)

	cronWorker.ProductStore = productStore

	cronWorker.Logger = applog

	cronWorker.UpdateProduct()

	applog.Info("Server started!")
	applog.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), server))
}
