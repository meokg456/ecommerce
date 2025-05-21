package kafkaservice

type OrderMessage struct {
	Id      string        `json:"id"`
	UserId  int           `json:"user_id"`
	Status  string        `json:"status"`
	Payment string        `json:"payment"`
	Paid    bool          `json:"paid"`
	Items   []ItemMessage `json:"items"`
}

type ItemMessage struct {
	ProductId string   `json:"product_id"`
	Types     []string `json:"types"`
	Quantity  int      `json:"quantity"`
}
