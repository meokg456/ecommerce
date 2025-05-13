package product

type Product struct {
	Id   string
	Name string
}

type Storage interface {
	GetProductById(id string) (*Product, error)
	AddProducts(products []Product) error
}
