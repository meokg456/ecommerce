package model

type GetProductByIdRequest struct {
	Id string `param:"id" validate:"required"`
}

type GetProductByIdResponse struct {
	Id           string   `json:"id"`
	Title        string   `json:"title"`
	Descriptions string   `json:"descriptions"`
	Category     string   `json:"category"`
	Images       []string `json:"images"`
	AdditionInfo any      `json:"addition_info"`
}
