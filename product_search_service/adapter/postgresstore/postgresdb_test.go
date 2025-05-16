package postgresstore_test

import (
	"testing"

	"github.com/meokg456/productsearchservice/adapter/testutil"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	dbName, dbUser, DbPassword := "db", "user", "pass"
	db := testutil.CreateConnection(t, dbName, dbUser, DbPassword)
	testutil.MigrateTestDatabase(t, db, "../../migrations")

	var current_user string
	err := db.Raw("SELECT current_user").Scan(&current_user).Error

	assert.NoError(t, err)
	assert.Equal(t, dbUser, current_user)
}
