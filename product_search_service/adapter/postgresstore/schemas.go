package postgresstore

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type ProductQuerySchema struct {
	Id           string
	Title        string
	Descriptions string
	Category     string
	Images       pq.StringArray `gorm:"type:text[]"`
	AdditionInfo map[string]any `gorm:"serializer:json"`
	MerchantId   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}
