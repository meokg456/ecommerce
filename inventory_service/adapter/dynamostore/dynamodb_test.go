package dynamostore_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/meokg456/inventoryservice/adapter/dynamostore"
	"github.com/meokg456/inventoryservice/adapter/testutil"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	db := testutil.CreateDynamoConnection(t, "us-east-1")
	dynamostore.MigrateDatabase(db)

	input := &dynamodb.ListTablesInput{}
	_, err := db.ListTables(context.Background(), input)

	assert.NoError(t, err)
}
