package product

import "github.com/meokg456/productservice/domain/common"

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
	GetProductsByMerchantId(merchantId int, page common.Page) ([]Product, string, error)
	GetProductById(id string) (*Product, error)
	AddProducts(products []Product) error
	AddProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(merchantId int, id string) error
}

type Broker interface {
	SendProductChange(p Product) error
}

func NewProductWithId(id string, title string, descriptions string, category string, images []string, additionInfo map[string]any) Product {
	return Product{
		Id:           id,
		Title:        title,
		Descriptions: descriptions,
		Category:     category,
		Images:       images,
		AdditionInfo: additionInfo,
	}
}
