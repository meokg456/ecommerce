package dynamostore_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/meokg456/productservice/adapter/dynamostore"
	"github.com/meokg456/productservice/adapter/testutil"
	"github.com/meokg456/productservice/dbconst"
	"github.com/meokg456/productservice/domain/product"
	"github.com/stretchr/testify/assert"
)

func TestProductStore(t *testing.T) {
	db := testutil.CreateDynamoConnection(t, "us-east-1")
	dynamostore.MigrateDatabase(db)

	store := dynamostore.NewProductStore(db)

	t.Run("Test Add products", func(t *testing.T) {
		products := testutil.Products

		err := store.AddProducts(products)

		assert.NoError(t, err)

		for _, p := range products {
			VerifyProduct(t, db, p)
		}
	})

	t.Run("Test Get product by id", func(t *testing.T) {
		products := testutil.Products

		err := store.AddProducts(products)

		assert.NoError(t, err)

		p, err := store.GetProductById("2")
		assert.NoError(t, err)

		assert.Equal(t, p.Id, "2")
		assert.Equal(t, p.Title, "Smartphone")
	})
}

func VerifyProduct(t *testing.T, client *dynamodb.Client, p product.Product) {
	t.Helper()

	data := dynamostore.ProductData{
		ID: p.Id,
	}
	marshalId, err := attributevalue.Marshal(data.ID)

	assert.NoError(t, err)

	output, err := client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(dbconst.ProductTableName),
		Key:       map[string]types.AttributeValue{"ID": marshalId},
	})

	assert.NoError(t, err)

	assert.NotEqual(t, len(output.Item), 0)

	err = attributevalue.UnmarshalMap(output.Item, &data)
	assert.NoError(t, err)

	assert.Equal(t, p.Id, data.ID)
	assert.Equal(t, p.Title, data.Title)
	assert.Equal(t, p.Descriptions, data.Descriptions)
	assert.Equal(t, p.Category, data.Category)
	assert.Equal(t, p.Images, data.Images)
	assert.Equal(t, p.AdditionInfo, data.AdditionInfo)
}
