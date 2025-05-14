package user_test

import (
	"testing"

	"github.com/meokg456/productmanagement/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestComparePassword(t *testing.T) {
	password := "123456"
	hashedPassword := "$2a$12$DGtYlBOVjO7Jd0ePCrOM/OJ7ICPkLQ/rZVfKKxL83BFWFTv36Up0m"

	t.Run("Compare return no error", func(t *testing.T) {
		assert.NoError(t, user.ComparePassword(hashedPassword, password))
	})

	t.Run("Compare return error", func(t *testing.T) {
		assert.Error(t, user.ComparePassword(hashedPassword, "wrongpassword"))
	})
}
