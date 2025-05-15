package httpserver_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meokg456/productmanagement/adapter/httpserver"
	"github.com/meokg456/productmanagement/adapter/httpserver/model"
	"github.com/meokg456/productmanagement/adapter/testutil"
	"github.com/meokg456/productmanagement/domain/product"
	"github.com/meokg456/productmanagement/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProductService struct {
	mock.Mock
}

func (p *ProductService) AddProduct(product *product.Product) error {
	args := p.Called(product)
	return args.Error(0)
}

func (p *ProductService) UpdateProduct(product *product.Product) error {
	args := p.Called(product)
	return args.Error(0)
}

func (p *ProductService) DeleteProduct(merchantId int, id string) error {
	args := p.Called(merchantId, id)
	return args.Error(0)
}

func TestAddProduct(t *testing.T) {

	server := httpserver.New(new(config.Config))
	mockStore := new(ProductService)
	server.ProductService = mockStore
	server.Logger = setupLogger(t)
	addProductRequest := model.AddProductRequest{
		Title:        "Laptop",
		Descriptions: "A high-performance laptop for work and gaming.",
		Category:     "Electronics",
		Images:       []string{"laptop1.jpg", "laptop2.jpg"},
		AdditionInfo: map[string]any{"brand": "BrandX", "warranty": "2 years"},
	}

	p := product.Product{
		Title:        addProductRequest.Title,
		Descriptions: addProductRequest.Descriptions,
		Category:     addProductRequest.Category,
		Images:       addProductRequest.Images,
		AdditionInfo: addProductRequest.AdditionInfo,
	}

	t.Run("Add product success", func(t *testing.T) {

		mockStore.On("AddProduct", &p).Return(nil).Once()

		data, err := json.Marshal(addProductRequest)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/products", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		expectedData := model.AddProductResponse{
			Title:        p.Title,
			Descriptions: p.Descriptions,
			Category:     p.Category,
			Images:       p.Images,
			AdditionInfo: p.AdditionInfo,
		}

		ctx := server.Router.NewContext(request, response)

		err = server.AddProduct(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, testutil.BuildSuccessBody(t, expectedData, http.StatusOK), response.Body.String())
		mockStore.AssertExpectations(t)
	})

	t.Run("Add product fail due to wrong request body format", func(t *testing.T) {

		data, err := json.Marshal(map[string]any{
			"wrong": "wrong",
		})
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/products", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		ctx := server.Router.NewContext(request, response)

		err = server.AddProduct(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, testutil.BuildErrorBody(t, http.StatusBadRequest, 0), response.Body.String())
	})

	t.Run("Register fail due to error when save user", func(t *testing.T) {
		mockStore.On("AddProduct", &p).Return(errors.New("Error")).Once()

		data, err := json.Marshal(addProductRequest)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/api/products", bytes.NewBuffer(data))
		request.Header.Set("Content-Type", "application/json")

		ctx := server.Router.NewContext(request, response)

		err = server.AddProduct(ctx)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, testutil.BuildErrorBody(t, http.StatusInternalServerError, 0), response.Body.String())
		mockStore.AssertExpectations(t)
	})
}
