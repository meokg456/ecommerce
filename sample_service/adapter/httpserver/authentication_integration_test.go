package httpserver_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meokg456/sampleservice/adapter/httpserver"
	"github.com/meokg456/sampleservice/adapter/httpserver/model"
	"github.com/meokg456/sampleservice/adapter/postgresstore"
	"github.com/meokg456/sampleservice/adapter/testutil"
	"github.com/meokg456/sampleservice/pkg/config"
	"github.com/stretchr/testify/assert"
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
