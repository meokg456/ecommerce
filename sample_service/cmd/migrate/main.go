package main

import (
	"log"

	"github.com/meokg456/sampleservice/adapter/postgresstore"
	"github.com/meokg456/sampleservice/pkg/config"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load env configs %v", err)
	}

	options := postgresstore.ParseFromConfig(config)
	db, err := postgresstore.NewConnection(options)
	if err != nil {
		log.Fatalf("Cannot connect to db %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	total, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Cannot execute migration %v", err)
	}

	log.Printf("%d applied", total)
}
