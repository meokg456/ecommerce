package product

type Product struct {
	Id           string
	Title        string
	Descriptions string
	Category     string
	Images       []string
	AdditionInfo any
}

type Storage interface {
	GetProductById(id string) (*Product, error)
	AddProducts(products []Product) error
}

func NewProductWithId(id string, title string, descriptions string, category string, images []string, additionInfo any) Product {
	return Product{
		Id:           id,
		Title:        title,
		Descriptions: descriptions,
		Category:     category,
		Images:       images,
		AdditionInfo: additionInfo,
	}
}
