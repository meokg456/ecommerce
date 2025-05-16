package model

type SearchProductsRequest struct {
	Keyword string `json:"keyword" validate:"required"`
	Page    int    `json:"page" validate:"required"`
	Limit   int    `json:"limit" validate:"required"`
}

type ProductResponse struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Descriptions string   `json:"descriptions"`
	Category     string   `json:"category"`
	Images       []string `json:"images"`
	AdditionInfo any      `json:"addition_info"`
	MerchantId   int      `json:"merchant_id"`
}
