package dynamostore

import (
	"context"
	"errors"

	"github.com/meokg456/ecommerce/utilities/dynamodbutils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/meokg456/productservice/dbconst"
	"github.com/meokg456/productservice/domain/product"
)

type ProductStore struct {
	client *dynamodb.Client
}

func NewProductStore(client *dynamodb.Client) *ProductStore {
	return &ProductStore{
		client: client,
	}
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
		Id:           data.ID,
		Title:        data.Title,
		Descriptions: data.Descriptions,
		Category:     data.Category,
		Images:       data.Images,
		AdditionInfo: data.AdditionInfo,
	}, nil
}

func (p *ProductStore) AddProducts(products []product.Product) error {
	var data []ProductData

	for _, p := range products {
		data = append(data, ProductData{
			ID:           p.Id,
			Title:        p.Title,
			Descriptions: p.Descriptions,
			Category:     p.Category,
			Images:       p.Images,
			AdditionInfo: p.AdditionInfo,
		})
	}

	batches, err := dynamodbutils.SplitDynamoBatchRequest(data, dbconst.ProductTableName)
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

func (p *ProductStore) AddProduct(product *product.Product) error {
	id := uuid.NewString()

	product.Id = id
	data := ProductData{
		ID:           id,
		Title:        product.Title,
		Descriptions: product.Descriptions,
		Category:     product.Category,
		Images:       product.Images,
		AdditionInfo: product.AdditionInfo,
	}

	av, err := attributevalue.MarshalMap(data)
	if err != nil {
		return err
	}

	_, err = p.client.PutItem(context.Background(), &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(dbconst.ProductTableName),
		ConditionExpression: aws.String("attribute_not_exists(ID)"),
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductStore) UpdateProduct(product *product.Product) error {
	data := ProductData{
		ID:           product.Id,
		Title:        product.Title,
		Descriptions: product.Descriptions,
		Category:     product.Category,
		Images:       product.Images,
		AdditionInfo: product.AdditionInfo,
	}

	av, err := attributevalue.MarshalMap(data)
	if err != nil {
		return err
	}

	_, err = p.client.PutItem(context.Background(), &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(dbconst.ProductTableName),
		ConditionExpression: aws.String("attribute_exists(ID)"),
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductStore) DeleteProduct(id string) error {
	data := ProductData{
		ID: id,
	}

	key, err := attributevalue.Marshal(data.ID)
	if err != nil {
		return err
	}

	_, err = p.client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		Key:                 map[string]types.AttributeValue{"ID": key},
		TableName:           aws.String(dbconst.ProductTableName),
		ConditionExpression: aws.String("attribute_exists(ID)"),
	})

	if err != nil {
		return err
	}

	return nil
}
