package postgresstore

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ProductQuerySchema struct {
	gorm.Model
	Id           string
	Title        string
	Descriptions string
	Category     string
	Images       pq.StringArray `gorm:"type:text[]"`
	AdditionInfo map[string]any `gorm:"serializer:json"`
	MerchantId   int
}
