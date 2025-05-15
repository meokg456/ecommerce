package model

type AddProductRequest struct {
	Title        string         `json:"title" validate:"required"`
	Descriptions string         `json:"descriptions" validate:"required"`
	Category     string         `json:"category" validate:"required"`
	Images       []string       `json:"images" validate:"required"`
	AdditionInfo map[string]any `json:"addition_info" validate:"required"`
}

type AddProductResponse struct {
	Id           string         `json:"id"`
	Title        string         `json:"title"`
	Descriptions string         `json:"descriptions"`
	Category     string         `json:"category"`
	Images       []string       `json:"images"`
	AdditionInfo map[string]any `json:"addition_info"`
	MerchantId   int            `json:"merchant_id"`
}

type UpdateProductRequest struct {
	Id           string         `param:"id" validate:"required"`
	Title        string         `json:"title" validate:"required"`
	Descriptions string         `json:"descriptions" validate:"required"`
	Category     string         `json:"category" validate:"required"`
	Images       []string       `json:"images" validate:"required"`
	AdditionInfo map[string]any `json:"addition_info" validate:"required"`
}

type UpdateProductResponse struct {
	Id           string         `json:"id"`
	Title        string         `json:"title"`
	Descriptions string         `json:"descriptions"`
	Category     string         `json:"category"`
	Images       []string       `json:"images"`
	AdditionInfo map[string]any `json:"addition_info"`
	MerchantId   int            `json:"merchant_id"`
}

type DeleteProductRequest struct {
	Id string `param:"id" validate:"required"`
}
