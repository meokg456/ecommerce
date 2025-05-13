package httpserver_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meokg456/userservice/adapter/httpserver"
	"github.com/meokg456/userservice/adapter/httpserver/model"
	"github.com/meokg456/userservice/adapter/postgresstore"
	"github.com/meokg456/userservice/adapter/testutil"
	"github.com/meokg456/userservice/domain/user"
	"github.com/meokg456/userservice/pkg/config"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestRegisterAPI(t *testing.T) {
	dbName, dbUser, DbPassword := "db", "user", "pass"
	db := testutil.CreateConnection(t, dbName, dbUser, DbPassword)
	testutil.MigrateTestDatabase(t, db, "../../migrations")

	server := httpserver.New(new(config.Config))
	server.UserStore = postgresstore.NewUserStore(db)
	server.Logger = setupLogger(t)

	registerUser := model.RegisterRequest{
		Username: "meokg456",
		Password: "123456",
		FullName: "Dung",
	}

	t.Run("Register success", func(t *testing.T) {
		data, err := json.Marshal(registerUser)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		expectedData := model.RegisterDataResponse{
			ID:       1,
			Username: registerUser.Username,
			FullName: registerUser.FullName,
		}

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, testutil.BuildSuccessBody(t, expectedData, http.StatusOK), response.Body.String())
	})
}

func TestLoginAPI(t *testing.T) {
	dbName, dbUser, DbPassword := "db", "user", "pass"
	db := testutil.CreateConnection(t, dbName, dbUser, DbPassword)
	testutil.MigrateTestDatabase(t, db, "../../migrations")

	jwtSecret := "secret"
	server := httpserver.New(&config.Config{
		JwtSecret: jwtSecret,
	})
	server.UserStore = postgresstore.NewUserStore(db)
	server.Logger = setupLogger(t)

	loginUser := model.LoginRequest{
		Username: "meokg456",
		Password: "123456",
	}

	t.Run("Login success", func(t *testing.T) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginUser.Password), 12)
		assert.NoError(t, err)

		u := user.NewUser("meokg456", string(hashedPassword), "Dung")
		err = server.UserStore.Register(&u)
		assert.NoError(t, err)

		data, err := json.Marshal(loginUser)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		assert.NoError(t, err)

		server.ServeHTTP(response, request)
		var responseBody map[string]any
		json.Unmarshal(response.Body.Bytes(), &responseBody)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Contains(t, responseBody["result"], "token")
	})
}
