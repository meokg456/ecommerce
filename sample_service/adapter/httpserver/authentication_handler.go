package httpserver

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meokg456/sampleservice/adapter/httpserver/model"
	"github.com/meokg456/sampleservice/domain/user"
	"go.uber.org/zap"
)

func (s *Server) RegisterAuthenticationRoute(router *echo.Group) {
	router.GET("/sample", s.Register)
}

func (s *Server) Register(c echo.Context) error {
	requestInfo := zap.String("request_id", s.requestID(c))
	var request model.RegisterRequest

	if err := c.Bind(&request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("register: error when parse request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	if err := c.Validate(request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("register: invalid request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	hashedPassword, err := user.HashPassword(request.Password)
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("register: hash password"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	u := user.NewUser(request.Username, hashedPassword, request.FullName)

	err = s.UserStore.Register(&u)
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("register user "+request.Username), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	return s.handleSuccess(c, model.RegisterDataResponse{
		ID:       u.ID,
		Username: u.Username,
		FullName: u.FullName,
	}, http.StatusOK)
}
