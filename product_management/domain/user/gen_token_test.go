package user_test

import (
	"testing"

	"github.com/meokg456/productmanagement/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestGenToken(t *testing.T) {
	token, err := user.GenToken(1, "secret")
	assert.NoError(t, err)
	assert.Contains(
		t,
		token,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
	)
}
