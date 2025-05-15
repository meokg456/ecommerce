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
	router.POST("/:id", s.UpdateProduct)
	router.DELETE("/:id", s.DeleteProduct)
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

func (s *Server) UpdateProduct(c echo.Context) error {
	requestInfo := zap.String("request_id", s.requestID(c))

	var request model.UpdateProductRequest
	if err := c.Bind(&request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("update product: error when parse request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	if err := c.Validate(request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("update product: invalid request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	p := product.NewProductWithId(
		request.Id,
		request.Title,
		request.Descriptions,
		request.Category,
		request.Images,
		request.AdditionInfo,
	)

	err := s.ProductService.UpdateProduct(&p)
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("update product: error when call grpc to update product"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	return s.handleSuccess(c, model.UpdateProductResponse{
		Id:           p.Id,
		Title:        p.Title,
		Descriptions: p.Descriptions,
		Category:     p.Category,
		Images:       p.Images,
		AdditionInfo: p.AdditionInfo,
	}, http.StatusOK)
}

func (s *Server) DeleteProduct(c echo.Context) error {
	requestInfo := zap.String("request_id", s.requestID(c))

	var request model.DeleteProductRequest
	if err := c.Bind(&request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("delete product: error when parse request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	if err := c.Validate(request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("delete product: invalid request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	err := s.ProductService.DeleteProduct(request.Id)
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("delete product: error when call grpc to delete product"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	return s.handleSuccess(c, request.Id, http.StatusOK)
}
