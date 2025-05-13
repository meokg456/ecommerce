package user_test

import (
	"testing"

	"github.com/meokg456/userservice/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "123456789"

	t.Run("Success with normal password", func(t *testing.T) {
		hashedPassword, err := user.HashPassword(password)
		assert.NoError(t, err)
		assert.NoError(t, user.ComparePassword(hashedPassword, password))
	})

	t.Run("Fail with too long password", func(t *testing.T) {
		longPassword := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
		hashedPassword, err := user.HashPassword(longPassword)
		assert.Error(t, err)
		assert.Equal(t, "", hashedPassword)
	})
}
