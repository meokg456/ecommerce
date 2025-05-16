package testutil

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/meokg456/inventoryservice/adapter/dynamostore"
	"github.com/stretchr/testify/assert"
)

func CreateDynamoConnection(t testing.TB, region string) *dynamodb.Client {
	ctx := context.Background()
	container := SetupDynamoDBContainer(t)

	port, _ := container.MappedPort(ctx, "8000")

	dynamodb, err := dynamostore.NewConnection(dynamostore.Options{
		Region:   "us-east-1",
		Endpoint: fmt.Sprintf("http://localhost:%d", port.Int()),
	})

	assert.NoError(t, err)

	return dynamodb
}
