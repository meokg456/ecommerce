package httpserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meokg456/sampleservice/adapter/httpserver"
	"github.com/meokg456/sampleservice/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestCheckHealthz(t *testing.T) {
	server := httpserver.New(new(config.Config))

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/healthz", nil)

	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Contains(t, response.Body.String(), "{\"messages\":\"Server is up and running\",\"status\":\"OK\"}")
}
