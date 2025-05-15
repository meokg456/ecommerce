package httpserver

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meokg456/productmanagement/adapter/httpserver/model"
	"github.com/meokg456/productmanagement/domain/common"
	"github.com/meokg456/productmanagement/domain/product"
	"go.uber.org/zap"
)

func (s *Server) RegisterProductRoute(router *echo.Group) {
	router.GET("", s.GetProducts)
	router.POST("", s.AddProduct)
	router.POST("/:id", s.UpdateProduct)
	router.DELETE("/:id", s.DeleteProduct)
}

func (s *Server) GetProducts(c echo.Context) error {
	requestInfo := zap.String("request_id", s.requestID(c))

	var request model.GetProductsByMerchantId
	if err := c.Bind(&request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("get products: error when parse request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	if err := c.Validate(request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("get products: invalid request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	userId := int(c.Get("user_id").(float64))

	products, lastKey, err := s.ProductService.GetProductsByMerchantId(userId, common.Page{
		Limit:         request.Limit,
		LastKeyOffset: request.LastKeyOffset,
	})
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("add product: error when call grpc to add product"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	var response []model.ProductResponse

	for _, p := range products {
		response = append(response, model.ProductResponse{
			Id:           p.Id,
			Title:        p.Title,
			Descriptions: p.Descriptions,
			Category:     p.Category,
			Images:       p.Images,
			AdditionInfo: p.AdditionInfo,
			MerchantId:   p.MerchantId,
		})
	}

	return s.handleSuccessWithPagination(c, http.StatusOK, response, lastKey, request.Limit, 0)
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

	userId := int(c.Get("user_id").(float64))

	p := product.NewProduct(
		request.Title,
		request.Descriptions,
		request.Category,
		request.Images,
		request.AdditionInfo,
		userId,
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
		MerchantId:   p.MerchantId,
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

	userId := int(c.Get("user_id").(float64))

	p := product.NewProductWithId(
		request.Id,
		request.Title,
		request.Descriptions,
		request.Category,
		request.Images,
		request.AdditionInfo,
		userId,
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
		MerchantId:   p.MerchantId,
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

	userId := int(c.Get("user_id").(float64))

	err := s.ProductService.DeleteProduct(userId, request.Id)
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("delete product: error when call grpc to delete product"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	return s.handleSuccess(c, request.Id, http.StatusOK)
}
