package model

type GetProductByIdRequest struct {
	Id int `param:"id" validate:"required"`
}
