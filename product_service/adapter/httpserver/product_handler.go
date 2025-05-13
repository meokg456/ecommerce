package httpserver

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meokg456/productservice/adapter/httpserver/model"
	"go.uber.org/zap"
)

func (s *Server) RegisterProductRoute(router *echo.Group) {
	router.GET("/:id", s.GetProductById)
}

func (s *Server) GetProductById(c echo.Context) error {
	requestInfo := zap.String("request_id", s.requestID(c))

	var request model.GetProductByIdRequest

	if err := c.Bind(&request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("get product by id: error when parse request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	if err := c.Validate(request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("get product by id: invalid request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	product, err := s.ProductStore.GetProductById(request.Id)
	if err != nil {
		s.Logger.Errorw(errors.Join(fmt.Errorf("get product by id: error when get product %s", request.Id), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	return s.handleSuccess(c, product, http.StatusOK)
}
