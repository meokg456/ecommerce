package httpserver

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/meokg456/productmanagement/domain/product"
	"github.com/meokg456/productmanagement/domain/user"
	"github.com/meokg456/productmanagement/pkg/config"
	"go.uber.org/zap"
)

type Server struct {
	Router *echo.Echo
	Config *config.Config
	Logger *zap.SugaredLogger

	UserStore      user.Storage
	ProductService product.Service
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New(config *config.Config) *Server {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	s := &Server{
		Router: e,
		Config: config,
	}

	s.RegisterGlobalMiddlewares()

	s.RegisterHealthCheck()

	apiGroup := s.Router.Group("/api")

	s.RegisterAuthenticationRoute(apiGroup)

	productGroup := apiGroup.Group("/products")

	s.RegisterProductRoute(productGroup)

	s.RegisterAuthMiddleware()

	return s
}

func (s *Server) RegisterAuthMiddleware() {
	skipPaths := []string{
		"/healthz",
		"/api/login",
		"/api/register",
		"/api/login-with-google",
	}

	s.Router.Use(echojwt.WithConfig(
		echojwt.Config{
			Skipper: func(c echo.Context) bool {
				return slices.Contains(skipPaths, c.Path())
			},
			SuccessHandler: func(c echo.Context) {
				user := c.Get("user").(*jwt.Token)
				claims := user.Claims.(jwt.MapClaims)
				c.Set("user_id", claims["user_id"])
			},
			SigningKey: []byte(s.Config.JwtSecret),
		},
	))
}

func (s *Server) RegisterGlobalMiddlewares() {
	s.Router.Use(middleware.Recover())
	s.Router.Use(middleware.Secure())
	s.Router.Use(middleware.RequestID())
	s.Router.Use(middleware.Gzip())

	// CORS
	if s.Config.AllowOrigins != "" {
		aos := strings.Split(s.Config.AllowOrigins, ",")
		s.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: aos,
		}))
	}
}

func (s *Server) RegisterHealthCheck() {
	s.Router.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":   http.StatusText(http.StatusOK),
			"messages": "Server is up and running",
		})
	})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) handleError(c echo.Context, status int, errCode int) error {

	messages := "Error!"
	if status == http.StatusBadRequest {
		messages = http.StatusText(status)
	}

	return c.JSON(status, map[string]string{
		"code":    strconv.Itoa(status) + fmt.Sprintf("%03d", errCode),
		"message": messages,
		"info":    messages,
	})
}

func (s *Server) handleSuccess(c echo.Context, data any, status int) error {

	return c.JSON(status, map[string]any{
		"code":    strconv.Itoa(status),
		"message": "OK",
		"result":  data,
	})
}

func (s *Server) handleSuccessWithPagination(c echo.Context, status int, data any, keyOffset string, limit int, total int) error {

	return c.JSON(status, map[string]any{
		"code":    strconv.Itoa(status),
		"message": "OK",
		"result": map[string]any{
			"data":       data,
			"key_offset": keyOffset,
			"limit":      limit,
			"total":      total,
		},
	})
}

func (s *Server) requestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}
