package user

import (
	"context"

	"google.golang.org/api/idtoken"
)

func ValidateGoogleIdToken(token string, clientId string) (*User, error) {
	payload, err := idtoken.Validate(context.Background(), token, clientId)
	if err != nil {
		return nil, err
	}

	return &User{
		Username: payload.Claims["email"].(string),
		Password: "",
		FullName: payload.Claims["name"].(string),
	}, nil
}
