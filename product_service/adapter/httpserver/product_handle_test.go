package httpserver_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meokg456/productservice/adapter/httpserver"
	"github.com/meokg456/productservice/adapter/httpserver/model"
	"github.com/meokg456/productservice/adapter/testutil"
	"github.com/meokg456/productservice/domain/product"
	"github.com/meokg456/productservice/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProductStore struct {
	mock.Mock
}

func (productStore *ProductStore) AddProducts(products []product.Product) error {
	args := productStore.Called(products)
	return args.Error(0)
}

func (productStore *ProductStore) GetProductById(id string) (*product.Product, error) {
	args := productStore.Called(id)
	return args.Get(0).(*product.Product), args.Error(1)
}

func TestGetProductById(t *testing.T) {

	server := httpserver.New(new(config.Config))
	mockStore := new(ProductStore)
	server.ProductStore = mockStore
	server.Logger = testutil.SetupLogger(t)
	request := model.GetProductByIdRequest{
		Id: "1",
	}

	t.Run("GetProductById success", func(t *testing.T) {
		p := testutil.Products[0]
		mockStore.On("GetProductById", request.Id).Return(&p, nil).Once()

		data, err := json.Marshal(request)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		expectedData := model.GetProductByIdResponse{
			Id:           p.Id,
			Title:        p.Title,
			Descriptions: p.Descriptions,
			Category:     p.Category,
			Images:       p.Images,
			AdditionInfo: p.AdditionInfo,
		}

		ctx := server.Router.NewContext(request, response)

		err = server.GetProductById(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, testutil.BuildSuccessBody(t, expectedData, http.StatusOK), response.Body.String())
		mockStore.AssertExpectations(t)
	})

	t.Run("GetProductById fail due to wrong request body format", func(t *testing.T) {

		data, err := json.Marshal(map[string]any{
			"ids": "wrong",
		})
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		ctx := server.Router.NewContext(request, response)

		err = server.GetProductById(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, testutil.BuildErrorBody(t, http.StatusBadRequest, 0), response.Body.String())
	})

	t.Run("GetProductById fail due to error when save user", func(t *testing.T) {
		var p *product.Product = nil
		mockStore.On("GetProductById", request.Id).Return(p, errors.New("Error")).Once()

		data, err := json.Marshal(request)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		ctx := server.Router.NewContext(request, response)

		err = server.GetProductById(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, testutil.BuildErrorBody(t, http.StatusInternalServerError, 0), response.Body.String())
		mockStore.AssertExpectations(t)
	})
}
