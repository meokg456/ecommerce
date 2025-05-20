package order

type Status string

const (
	Pending   Status = "Pending"
	Active    Status = "Active"
	Shipping  Status = "Shipping"
	Done      Status = "Done"
	Cancelled Status = "Cancelled"
)
