package main

import (
	"log"

	"github.com/meokg456/inventoryservice/adapter/dynamostore"
	"github.com/meokg456/inventoryservice/pkg/config"
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

	dynamostore.MigrateDatabase(db)

	log.Println("Migrated done")
}
