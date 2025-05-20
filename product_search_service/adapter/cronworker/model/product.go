package model

type ProductMessage struct {
	Id           string         `json:"id"`
	Title        string         `json:"title"`
	Descriptions string         `json:"descriptions"`
	Category     string         `json:"category"`
	Images       []string       `json:"images"`
	AdditionInfo map[string]any `json:"addition_info"`
	MerchantId   int            `json:"merchant_id"`
}
