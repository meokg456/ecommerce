package dynamostore

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/meokg456/ecommerce/utilities/dynamodbutils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/meokg456/productservice/dbconst"
	"github.com/meokg456/productservice/domain/common"
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

func (p *ProductStore) GetProductsByMerchantId(merchantId int, page common.Page) ([]product.Product, error) {
	var products []product.Product

	merchantIdString := strconv.Itoa(merchantId)

	output, err := p.client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(dbconst.ProductTableName),
		KeyConditionExpression: aws.String("MerchantId = :merchantId"),
		IndexName:              aws.String(dbconst.ProductMerchantIdIndexName),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":merchantId": &types.AttributeValueMemberN{Value: merchantIdString},
		},
		Limit: aws.Int32(int32(page.Limit)),
		ExclusiveStartKey: map[string]types.AttributeValue{
			dbconst.ProductPK:             &types.AttributeValueMemberS{Value: page.LastKeyOffset},
			dbconst.ProductMerchantIdName: &types.AttributeValueMemberN{Value: merchantIdString},
		},
	})

	if err != nil {
		return nil, err
	}

	fmt.Println(output.LastEvaluatedKey["ID"])
	fmt.Println(output.LastEvaluatedKey["MerchantId"])

	err = attributevalue.UnmarshalListOfMaps(output.Items, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
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
			MerchantId:   p.MerchantId,
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
		MerchantId:   product.MerchantId,
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
		MerchantId:   product.MerchantId,
	}

	av, err := attributevalue.MarshalMap(data)
	if err != nil {
		return err
	}

	_, err = p.client.PutItem(context.Background(), &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(dbconst.ProductTableName),
		ConditionExpression: aws.String("attribute_exists(ID) AND MerchantId = :merchantId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":merchantId": &types.AttributeValueMemberN{Value: strconv.Itoa(product.MerchantId)},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *ProductStore) DeleteProduct(merchantId int, id string) error {
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
		ConditionExpression: aws.String("attribute_exists(ID) AND MerchantId = :merchantId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":merchantId": &types.AttributeValueMemberN{Value: strconv.Itoa(merchantId)},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
