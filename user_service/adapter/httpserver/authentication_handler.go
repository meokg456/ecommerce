package httpserver

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meokg456/userservice/adapter/httpserver/model"
	"github.com/meokg456/userservice/domain/user"
	"go.uber.org/zap"
)

func (s *Server) RegisterAuthenticationRoute(router *echo.Group) {
	router.POST("/register", s.Register)
	router.POST("/login", s.Login)
	router.POST("/google-login", s.LoginWithGoogle)
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

func (s *Server) Login(c echo.Context) error {
	requestInfo := zap.String("request_id", s.requestID(c))
	var request model.LoginRequest

	if err := c.Bind(&request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("login: error when parse request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	if err := c.Validate(request); err != nil {
		s.Logger.Errorw(errors.Join(errors.New("login: invalid request body"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusBadRequest, 0)
	}

	u, err := s.UserStore.GetUserByUsername(request.Username)
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("login: Get user "+request.Username), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 1)
	}

	if user.ComparePassword(u.Password, request.Password) != nil {
		s.Logger.Errorw(errors.Join(errors.New("login: compare hash password "+request.Username), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 2)
	}

	token, err := user.GenToken(u.ID, s.Config.JwtSecret)
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("login: gen token "+request.Username), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	return s.handleSuccess(c, map[string]string{
		"token": token,
	}, http.StatusOK)
}

func (s *Server) LoginWithGoogle(c echo.Context) error {
	requestInfo := zap.String("request_id", s.requestID(c))

	csrfCookie, err := c.Cookie("g_csrf_token")
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("g_csrf_token cookie not exists"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	csrf := c.FormValue("g_csrf_token")

	if csrfCookie.Value != csrf {
		s.Logger.Errorw(errors.Join(errors.New("g_csrf_token is invalid"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	token := c.FormValue("credential")

	u, err := user.ValidateGoogleIdToken(token, s.Config.ClientId)
	if err != nil {
		s.Logger.Errorw(errors.Join(errors.New("id token invalid"), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	var userId int
	existingUser, err := s.UserStore.GetUserByUsername(u.Username)
	if err != nil {
		// not exist

		err = s.UserStore.Register(u)
		if err != nil {
			s.Logger.Errorw(errors.Join(errors.New("cannot register user "+u.Username), err).Error(), requestInfo)
			return s.handleError(c, http.StatusInternalServerError, 0)
		}
		userId = u.ID
	} else {
		userId = existingUser.ID
	}

	tokenString, err := user.GenToken(userId, s.Config.JwtSecret)
	if err != nil {
		s.Logger.Errorw(errors.Join(fmt.Errorf("cannot gen token %d", userId), err).Error(), requestInfo)
		return s.handleError(c, http.StatusInternalServerError, 0)
	}

	return s.handleSuccess(c, map[string]string{
		"token": tokenString,
	}, http.StatusOK)
}
