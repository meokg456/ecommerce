package httpserver

import "github.com/labstack/echo/v4"

func (s *Server) RegisterInventoryRoute(router *echo.Group) {
	router.GET("", s.GetProducts)
	router.POST("", s.AddProduct)
	router.POST("/:id", s.UpdateProduct)
	router.DELETE("/:id", s.DeleteProduct)
}
