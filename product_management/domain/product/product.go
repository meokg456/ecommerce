package product

import "github.com/meokg456/productmanagement/domain/common"

type Product struct {
	Id           string
	Title        string
	Descriptions string
	Category     string
	Images       []string
	AdditionInfo map[string]any
	MerchantId   int
}

type Service interface {
	GetProductsByMerchantId(merchantId int, page common.Page) ([]Product, string, error)
	AddProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(merchantId int, int string) error
}

func NewProduct(title string, descriptions string, category string, images []string, additionInfo map[string]any, merchantId int) Product {
	return Product{
		Title:        title,
		Descriptions: descriptions,
		Category:     category,
		Images:       images,
		AdditionInfo: additionInfo,
		MerchantId:   merchantId,
	}
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
