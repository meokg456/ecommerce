package httpserver

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meokg456/productsearchservice/adapter/httpserver/model"
	"github.com/meokg456/productsearchservice/domain/common"
	"go.uber.org/zap"
)

func (s *Server) RegisterProductRoutes(router *echo.Group) {
	router.GET("/search", s.SearchProducts)
}

func (s *Server) SearchProducts(c echo.Context) error {
	requestInfo := zap.String("request_id", s.requestID(c))

	var request model.SearchProductsRequest
	if err := c.Bind(&request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("search products: error when parse request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	if err := c.Validate(request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("search products: invalid request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	products, err := s.ProductStore.SearchProducts(request.Keyword, common.Page{
		Page:  request.Page,
		Limit: request.Limit,
	})
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("search products: error when search products"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	return s.handleSuccessWithPagination(c, http.StatusOK, products, request.Page, request.Limit, 0)
}
