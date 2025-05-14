package httpserver

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meokg456/productmanagement/adapter/httpserver/model"
	"github.com/meokg456/productmanagement/domain/product"
	"go.uber.org/zap"
)

func (s *Server) RegisterProductRoute(router *echo.Group) {
	router.POST("", s.AddProduct)
}

func (s *Server) AddProduct(c echo.Context) error {
	requestInfo := zap.String("request_id", s.requestID(c))

	var request model.AddProductRequest
	if err := c.Bind(&request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("add product: error when parse request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	if err := c.Validate(request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("add product: invalid request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	p := product.NewProduct(
		request.Title,
		request.Descriptions,
		request.Category,
		request.Images,
		request.AdditionInfo,
	)

	err := s.ProductService.AddProduct(&p)
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("add product: error when call grpc to add product"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	return s.handleSuccess(c, model.AddProductResponse{
		Id:           p.Id,
		Title:        p.Title,
		Descriptions: p.Descriptions,
		Category:     p.Category,
		Images:       p.Images,
		AdditionInfo: p.AdditionInfo,
	}, http.StatusOK)
}
