package httpserver

import (
	"errors"
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
		s.Logger.Errorw(errors.Join(errors.New("register: error when parse request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	if err := c.Validate(request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("register: invalid request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	return s.handleSuccess(c, nil, http.StatusOK)
}
