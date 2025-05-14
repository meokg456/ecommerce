package product

type Product struct {
	Id           string
	Title        string
	Descriptions string
	Category     string
	Images       []string
	AdditionInfo map[string]any
}

type Service interface {
	AddProduct(product *Product) error
}

func NewProduct(title string, descriptions string, category string, images []string, additionInfo map[string]any) Product {
	return Product{
		Title:        title,
		Descriptions: descriptions,
		Category:     category,
		Images:       images,
		AdditionInfo: additionInfo,
	}
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
