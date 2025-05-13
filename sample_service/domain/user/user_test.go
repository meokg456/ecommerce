package user_test

import (
	"testing"

	"github.com/meokg456/sampleservice/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := user.NewUser("meokg456", "123456", "Dung")
	assert.Equal(t, user.Username, "meokg456")
	assert.Equal(t, user.Password, "123456")
	assert.Equal(t, user.FullName, "Dung")
}

func TestNewUserWithId(t *testing.T) {
	user := user.NewUserWithId(1, "meokg456", "123456", "Dung")
	assert.Equal(t, user.ID, 1)
	assert.Equal(t, user.Username, "meokg456")
	assert.Equal(t, user.Password, "123456")
	assert.Equal(t, user.FullName, "Dung")
}
