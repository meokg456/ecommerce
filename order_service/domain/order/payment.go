package order

type Payment string

const (
	COD     Payment = "COD"
	Banking Payment = "Banking"
	Paypal  Payment = "Paypal"
)
