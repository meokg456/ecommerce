package postgresstore

import (
	"gorm.io/gorm"
)

type UserQuerySchema struct {
	gorm.Model
	Username string
	Password string
	FullName string
}
