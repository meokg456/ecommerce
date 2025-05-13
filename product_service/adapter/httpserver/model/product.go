package model

type GetProductByIdRequest struct {
	Id string `param:"id" validate:"required"`
}
