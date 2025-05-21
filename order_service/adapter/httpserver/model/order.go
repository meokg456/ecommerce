package model

type OrderRequest struct {
	UserId  int    `json:"user_id" validate:"required"`
	Payment string `json:"payment" validate:"required"`
	Paid    bool   `json:"paid" validate:"required"`
	Items   []Item `json:"items" validate:"required"`
}

type Item struct {
	ProductId string   `json:"product_id" validate:"required"`
	Types     []string `json:"types" validate:"required"`
	Quantity  int      `json:"quantity" validate:"required"`
}

type OrderResponse struct {
	Id      string `json:"id"`
	UserId  int    `json:"user_id"`
	Payment string `json:"payment"`
	Status  string `json:"status"`
	Paid    bool   `json:"paid"`
	Items   []Item `json:"items"`
}
