package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenToken(userId int, jwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"iss":     "meokg456",
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(jwtSecret))
}
