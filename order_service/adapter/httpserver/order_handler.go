package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) RegisterOrderRoutes(router *echo.Group) {
	router.POST("", s.Order)
}

func (s *Server) Order(c echo.Context) error {
	return s.handleSuccess(c, nil, http.StatusOK)
}
