package dynamostore

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/meokg456/productservice/dbconst"
	"github.com/meokg456/productservice/domain/product"
	"github.com/meokg456/productservice/pkg/utils"
)

type ProductStore struct {
	client *dynamodb.Client
}

func (p *ProductStore) GetProductById(id string) (*product.Product, error) {
	data := ProductData{
		ID: id,
	}
	marshalId, err := attributevalue.Marshal(data.ID)

	if err != nil {
		return nil, err
	}

	output, err := p.client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(dbconst.ProductTableName),
		Key:       map[string]types.AttributeValue{"ID": marshalId},
	})

	if err != nil {
		return nil, err
	}

	if len(output.Item) == 0 {
		return nil, errors.New("item not found")
	}

	err = attributevalue.UnmarshalMap(output.Item, &data)
	if err != nil {
		return nil, err
	}

	return &product.Product{
		Id:   data.ID,
		Name: data.Name,
	}, nil
}

func (p *ProductStore) AddProducts(products []product.Product) error {
	var data []ProductData

	for _, p := range products {
		data = append(data, ProductData{
			ID:   p.Id,
			Name: p.Name,
		})
	}

	batches, err := utils.SplitDynamoBatchRequest(data, dbconst.ProductTableName)
	if err != nil {
		return err
	}

	for _, batch := range batches {
		_, err := p.client.BatchWriteItem(context.Background(), &batch)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewProductStore(client *dynamodb.Client) *ProductStore {
	return &ProductStore{
		client: client,
	}
}
