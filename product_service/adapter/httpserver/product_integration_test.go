package httpserver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meokg456/productservice/adapter/dynamostore"
	"github.com/meokg456/productservice/adapter/httpserver"
	"github.com/meokg456/productservice/adapter/httpserver/model"
	"github.com/meokg456/productservice/adapter/testutil"
	"github.com/meokg456/productservice/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestGetBookByIdAPI(t *testing.T) {
	dynamoDB := testutil.CreateDynamoConnection(t, "us-east-1")
	dynamostore.MigrateDatabase(dynamoDB)

	server := httpserver.New(new(config.Config))
	server.ProductStore = dynamostore.NewProductStore(dynamoDB)
	server.Logger = testutil.SetupLogger(t)

	t.Run("Get book by id success", func(t *testing.T) {
		products := testutil.Products

		server.ProductStore.AddProducts(products)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/api/products/1", nil)
		request.Header.Set("Content-Type", "application/json")

		p := products[0]
		expectedData := model.GetProductByIdResponse{
			Id:           p.Id,
			Title:        p.Title,
			Descriptions: p.Descriptions,
			Category:     p.Category,
			Images:       p.Images,
			AdditionInfo: p.AdditionInfo,
		}

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, testutil.BuildSuccessBody(t, expectedData, http.StatusOK), response.Body.String())
	})
}
