package postgresstore

import (
	"github.com/meokg456/productsearchservice/domain/common"
	"github.com/meokg456/productsearchservice/domain/product"
	"gorm.io/gorm"
)

type ProductStore struct {
	db *gorm.DB
}

func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{
		db: db,
	}
}

func (p *ProductStore) SearchProducts(keyword string, page common.Page) ([]product.Product, error) {
	var products []product.Product

	return products, nil
}
