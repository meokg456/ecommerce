package postgresstore

import (
	"gorm.io/gorm"
)

type UserQuerySchema struct {
	gorm.Model
	Username string
	Password string
	FullName string
	Avatar   string
}

type BookQuerySchema struct {
	gorm.Model
	Name       string
	Content    string
	Pages      int
	CategoryId int
	Category   CategoryQuerySchema
}

type CategoryQuerySchema struct {
	gorm.Model
	Name string
}
