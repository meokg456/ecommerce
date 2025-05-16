package testutil

import (
	"context"
	"testing"

	"github.com/meokg456/productsearchservice/adapter/postgresstore"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func MigrateTestDatabase(t testing.TB, db *gorm.DB, migrationPath string) {
	t.Helper()

	migrations := &migrate.FileMigrationSource{
		Dir: migrationPath,
	}

	sqlDB, err := db.DB()
	assert.NoError(t, err)

	_, err = migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	assert.NoError(t, err)
}

func CreateConnection(t testing.TB, dbName, dbUser, dbPassword string) *gorm.DB {
	cont := SetupPostgresContainer(t, dbName, dbUser, dbPassword)
	host, _ := cont.Host(context.Background())
	port, _ := cont.MappedPort(context.Background(), "5432")

	db, err := postgresstore.NewConnection(postgresstore.Options{
		DBName:   dbName,
		User:     dbUser,
		Password: dbPassword,
		Host:     host,
		Port:     port.Port(),
	})

	assert.NoError(t, err)

	return db
}
