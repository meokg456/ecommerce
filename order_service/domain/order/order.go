package order

type Payment string

type Order struct {
	Id      int
	UserId  int
	Status  Status
	Payment Payment
	Paid    bool
	Items   []Item
}

type Broker interface {
	SaveOrder(order Order) error
}
