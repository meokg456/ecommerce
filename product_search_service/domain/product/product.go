package product

import "github.com/meokg456/productsearchservice/domain/common"

type Product struct {
	Id           string
	Title        string
	Descriptions string
	Category     string
	Images       []string
	AdditionInfo map[string]any
	MerchantId   int
}

type Storage interface {
	SearchProducts(keyword string, page common.Page) ([]Product, error)
	SaveProduct(p Product) error
}

func NewProductWithId(id string, title string, descriptions string, category string, images []string, additionInfo map[string]any, merchantId int) Product {
	return Product{
		Id:           id,
		Title:        title,
		Descriptions: descriptions,
		Category:     category,
		Images:       images,
		AdditionInfo: additionInfo,
		MerchantId:   merchantId,
	}
}
