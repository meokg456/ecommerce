package product

type Product struct {
	Id   int
	Name string
}

type Storage interface {
	GetProductById(id int) (*Product, error)
}
