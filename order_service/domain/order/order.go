package order

type Order struct {
	Id      string
	UserId  int
	Status  Status
	Payment Payment
	Paid    bool
	Items   []Item
}

type Broker interface {
	SaveOrder(order *Order) error
}
