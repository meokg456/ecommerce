package httpserver

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meokg456/orderservice/adapter/httpserver/model"
	"github.com/meokg456/orderservice/domain/order"
	"go.uber.org/zap"
)

func (s *Server) RegisterOrderRoutes(router *echo.Group) {
	router.POST("", s.Order)
}

func (s *Server) Order(c echo.Context) error {
	requestInfo := zap.String("request_id", s.requestID(c))

	var request model.OrderRequest

	if err := c.Bind(&request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("order: error when parse request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	if err := c.Validate(request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("order: invalid request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	items := []order.Item{}
	for _, item := range request.Items {
		items = append(items, order.Item{
			ProductId: item.ProductId,
			Types:     item.Types,
			Quantity:  item.Quantity,
		})
	}

	order := order.Order{
		UserId:  request.UserId,
		Payment: order.Payment(request.Payment),
		Paid:    request.Paid,
		Items:   items,
	}

	err := s.OrderBroker.SaveOrder(&order)
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("order: cannot save order"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	return s.handleSuccess(c, model.OrderResponse{
		Id:      order.Id,
		UserId:  order.UserId,
		Payment: string(order.Payment),
		Status:  string(order.Status),
		Paid:    order.Paid,
		Items:   request.Items,
	}, http.StatusOK)
}
