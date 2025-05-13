package main

import (
	"log"

	"github.com/meokg456/productservice/adapter/dynamostore"
	"github.com/meokg456/productservice/domain/product"
	"github.com/meokg456/productservice/pkg/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load env configs %v", err)
	}

	options := dynamostore.ParseFromConfig(config)
	db, err := dynamostore.NewConnection(options)
	if err != nil {
		log.Fatalf("Cannot connect to db %v", err)
	}

	products := []product.Product{
		{Id: "1", Name: "Laptop"},
		{Id: "2", Name: "Smartphone"},
		{Id: "3", Name: "Headphones"},
		{Id: "4", Name: "Keyboard"},
		{Id: "5", Name: "Mouse"},
		{Id: "6", Name: "Monitor"},
		{Id: "7", Name: "Tablet"},
		{Id: "8", Name: "Smartwatch"},
		{Id: "9", Name: "Printer"},
		{Id: "10", Name: "Camera"},
	}

	store := dynamostore.NewProductStore(db)

	err = store.AddProducts(products)

	if err != nil {
		log.Fatalf("Cannot seed data %v", err)
	}
}
