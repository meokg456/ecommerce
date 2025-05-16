package inventory

type Inventory struct {
	ProductId string
	Types     []string
	Quantity  int
}

type Storage interface {
	SaveInventory(inventory Inventory) error
}

func NewInventory(productId string, types []string, quantity int) Inventory {
	return Inventory{
		ProductId: productId,
		Types:     types,
		Quantity:  quantity,
	}
}
