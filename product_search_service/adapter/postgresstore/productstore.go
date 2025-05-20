package postgresstore

import (
	"strings"

	"github.com/meokg456/productsearchservice/dbconst"
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
	products := []product.Product{}

	var productsData []ProductQuerySchema

	tsquery := strings.Join(strings.Split(keyword, " "), " & ")

	result := p.db.
		Table(dbconst.ProductTableName).
		Select("*, ts_rank_cd(tsv, to_tsquery(?)) AS rank", tsquery).
		Where("tsv @@ to_tsquery(?)", tsquery).
		Offset((page.Page - 1) * page.Limit).
		Limit(page.Limit).
		Order("rank DESC").
		Find(&productsData)

	for _, data := range productsData {
		products = append(products, product.NewProductWithId(
			data.Id, data.Title, data.Descriptions, data.Category, data.Images, data.AdditionInfo, data.MerchantId))
	}

	return products, result.Error
}

func (p *ProductStore) SaveProduct(pro product.Product) error {
	data := ProductQuerySchema{
		Id:           pro.Id,
		Title:        pro.Title,
		Descriptions: pro.Descriptions,
		Category:     pro.Category,
		Images:       pro.Images,
		AdditionInfo: pro.AdditionInfo,
		MerchantId:   pro.MerchantId,
	}
	result := p.db.Table(dbconst.ProductTableName).Save(&data)

	return result.Error
}
