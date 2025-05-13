package postgresstore_test

import (
	"testing"

	"github.com/meokg456/sampleservice/adapter/postgresstore"
	"github.com/meokg456/sampleservice/adapter/testutil"
	"github.com/meokg456/sampleservice/dbconst"
	"github.com/meokg456/sampleservice/domain/user"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserStore(t *testing.T) {
	dbName, dbUser, DbPassword := "db", "user", "pass"
	db := testutil.CreateConnection(t, dbName, dbUser, DbPassword)
	testutil.MigrateTestDatabase(t, db, "../../migrations")

	store := postgresstore.NewUserStore(db)

	t.Run("Test Register", func(t *testing.T) {
		user := user.NewUser("meokg456", "hashed password", "Dung")
		err := store.Register(&user)

		assert.NoError(t, err)
		VerifyRegisteredUser(t, user.Username, db)
	})
}

func VerifyRegisteredUser(t testing.TB, username string, db *gorm.DB) {
	t.Helper()

	var user postgresstore.UserQuerySchema
	result := db.Table(dbconst.UserTableName).First(&user, "username = ?", username)

	assert.NoError(t, result.Error)
}
