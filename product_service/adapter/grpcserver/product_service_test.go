package grpcserver_test

import (
	"context"
	"errors"
	"testing"

	pb "github.com/meokg456/ecommerce/proto/product"
	"github.com/meokg456/productservice/adapter/grpcserver"
	"github.com/meokg456/productservice/adapter/testutil"
	"github.com/meokg456/productservice/domain/common"
	"github.com/meokg456/productservice/domain/product"
	"github.com/meokg456/productservice/pkg/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/structpb"
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

func (productStore *ProductStore) AddProduct(product *product.Product) error {
	args := productStore.Called(product)
	return args.Error(0)
}

func (productStore *ProductStore) UpdateProduct(product *product.Product) error {
	args := productStore.Called(product)
	return args.Error(0)
}

func (productStore *ProductStore) DeleteProduct(merchantId int, id string) error {
	args := productStore.Called(merchantId, id)
	return args.Error(0)
}

func (productStore *ProductStore) GetProductsByMerchantId(merchantId int, page common.Page) ([]product.Product, string, error) {
	args := productStore.Called(merchantId, page)
	return args.Get(0).([]product.Product), args.String(1), args.Error(2)
}

func TestAddProduct(t *testing.T) {
	mockStore := new(ProductStore)

	grpcService := grpcserver.New(&config.Config{})
	grpcService.Logger = testutil.SetupLogger(t)
	grpcService.ProductStore = mockStore

	p := product.Product{
		Title:        "Laptop",
		Descriptions: "A high-performance laptop for work and gaming.",
		Category:     "Electronics",
		Images:       []string{"laptop1.jpg", "laptop2.jpg"},
		AdditionInfo: map[string]any{"brand": "BrandX", "warranty": "2 years"},
	}

	t.Run("Add product success", func(t *testing.T) {
		additionInfo, err := structpb.NewStruct(p.AdditionInfo)
		assert.NoError(t, err)

		request := pb.Product{
			Title:        p.Title,
			Descriptions: p.Descriptions,
			Category:     p.Category,
			Images:       p.Images,
			AdditionInfo: additionInfo,
		}

		mockStore.On("AddProduct", &p).Return(nil).Once()

		response, err := grpcService.AddProduct(context.Background(), &request)

		assert.NoError(t, err)
		assert.Equal(t, request.Id, response.Id)
		assert.Equal(t, request.Title, response.Title)
		assert.Equal(t, request.Descriptions, response.Descriptions)
		assert.Equal(t, request.Category, response.Category)
		assert.Equal(t, request.Images, response.Images)
	})

	t.Run("Add product fail", func(t *testing.T) {
		additionInfo, err := structpb.NewStruct(p.AdditionInfo)
		assert.NoError(t, err)

		request := pb.Product{
			Title:        p.Title,
			Descriptions: p.Descriptions,
			Category:     p.Category,
			Images:       p.Images,
			AdditionInfo: additionInfo,
		}

		mockStore.On("AddProduct", &p).Return(errors.New("failed")).Once()

		response, err := grpcService.AddProduct(context.Background(), &request)

		assert.Error(t, err)
		assert.Nil(t, response)
	})
}
