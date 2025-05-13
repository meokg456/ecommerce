package dynamostore

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/meokg456/productservice/domain/product"
)

type ProductStore struct {
	client *dynamodb.Client
}

func (p *ProductStore) GetProductById(id int) (*product.Product, error) {
	return nil, nil
}

func NewProductStore(client *dynamodb.Client) *ProductStore {
	return &ProductStore{
		client: client,
	}
}
