package httpserver_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meokg456/userservice/adapter/httpserver"
	"github.com/meokg456/userservice/adapter/httpserver/model"
	"github.com/meokg456/userservice/adapter/testutil"
	"github.com/meokg456/userservice/domain/user"
	"github.com/meokg456/userservice/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	mock.Mock
}

func (userStore *UserStore) Register(user *user.User) error {
	args := userStore.Called(user)
	return args.Error(0)
}

func (userStore *UserStore) GetUserByUsername(username string) (*user.User, error) {
	args := userStore.Called(username)
	return args.Get(0).(*user.User), args.Error(1)
}

func (userStore *UserStore) GetUserById(id int) (*user.User, error) {
	args := userStore.Called(id)
	return args.Get(0).(*user.User), args.Error(1)
}

func (userStore *UserStore) CheckIfUserExist(id int) error {
	args := userStore.Called(id)
	return args.Error(0)
}

func (userStore *UserStore) UpdateProfile(user *user.User) error {
	args := userStore.Called(user)
	return args.Error(0)
}

func TestRegister(t *testing.T) {

	server := httpserver.New(new(config.Config))
	mockStore := new(UserStore)
	server.UserStore = mockStore
	server.Logger = setupLogger(t)
	registerUser := model.RegisterRequest{
		Username: "meokg456",
		Password: "123456",
		FullName: "Dung",
	}

	t.Run("Register success", func(t *testing.T) {
		mockStore.On("Register", mock.Anything).Return(nil).Once()

		data, err := json.Marshal(registerUser)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		expectedData := model.RegisterDataResponse{
			ID:       0,
			Username: registerUser.Username,
			FullName: registerUser.FullName,
		}

		ctx := server.Router.NewContext(request, response)

		err = server.Register(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, testutil.BuildSuccessBody(t, expectedData, http.StatusOK), response.Body.String())
		mockStore.AssertExpectations(t)
	})

	t.Run("Register fail due to wrong request body format", func(t *testing.T) {

		data, err := json.Marshal(map[string]any{
			"user": "user",
			"pass": "pass",
			"full": "full",
		})
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		ctx := server.Router.NewContext(request, response)

		err = server.Register(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, testutil.BuildErrorBody(t, http.StatusBadRequest, 0), response.Body.String())
	})

	t.Run("Register fail due to error when save user", func(t *testing.T) {
		mockStore.On("Register", mock.Anything).Return(errors.New("Error")).Once()

		data, err := json.Marshal(registerUser)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		ctx := server.Router.NewContext(request, response)

		err = server.Register(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, testutil.BuildErrorBody(t, http.StatusInternalServerError, 0), response.Body.String())
		mockStore.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	jwtSecret := "secret"
	server := httpserver.New(&config.Config{
		JwtSecret: jwtSecret,
	})
	mockStore := new(UserStore)
	server.UserStore = mockStore
	server.Logger = setupLogger(t)
	loginUser := model.LoginRequest{
		Username: "meokg456",
		Password: "123456",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginUser.Password), 12)
	assert.NoError(t, err)

	u := user.NewUser("meokg456", string(hashedPassword), "Dung")

	t.Run("Login success", func(t *testing.T) {
		mockStore.On("GetUserByUsername", loginUser.Username).Return(&u, nil).Once()

		data, err := json.Marshal(loginUser)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		assert.NoError(t, err)

		ctx := server.Router.NewContext(request, response)

		err = server.Login(ctx)

		var responseBody map[string]any
		json.Unmarshal(response.Body.Bytes(), &responseBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, responseBody["result"], "token")
		mockStore.AssertExpectations(t)
	})

	t.Run("Login fail due to wrong request body format", func(t *testing.T) {

		data, err := json.Marshal(map[string]any{
			"user": "user",
			"pass": "pass",
		})
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		ctx := server.Router.NewContext(request, response)

		err = server.Login(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, testutil.BuildErrorBody(t, http.StatusBadRequest, 0), response.Body.String())
	})

	t.Run("Login fail due to error when get user", func(t *testing.T) {
		var result *user.User = nil
		mockStore.On("GetUserByUsername", loginUser.Username).Return(result, errors.New("Error")).Once()

		data, err := json.Marshal(loginUser)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		ctx := server.Router.NewContext(request, response)

		err = server.Login(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, testutil.BuildErrorBody(t, http.StatusInternalServerError, 1), response.Body.String())
		mockStore.AssertExpectations(t)
	})

	t.Run("Login fail due to wrong password", func(t *testing.T) {
		wrongPassUser := u
		wrongPassUser.Password = "wrong"
		mockStore.On("GetUserByUsername", loginUser.Username).Return(&wrongPassUser, nil).Once()

		data, err := json.Marshal(loginUser)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		assert.NoError(t, err)

		ctx := server.Router.NewContext(request, response)

		err = server.Login(ctx)

		var responseBody map[string]any
		json.Unmarshal(response.Body.Bytes(), &responseBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, testutil.BuildErrorBody(t, http.StatusInternalServerError, 2), response.Body.String())
		mockStore.AssertExpectations(t)
	})
}

func setupLogger(t *testing.T) *zap.SugaredLogger {
	t.Helper()

	logger, err := zap.NewProduction()
	defer logger.Sync()
	assert.NoError(t, err)
	return logger.Sugar()
}
