package testutil

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/dynamodb"
)

func SetupDynamoDBContainer(t testing.TB) testcontainers.Container {
	ctx := context.Background()

	ctr, err := dynamodb.Run(ctx, "amazon/dynamodb-local:1.19.0")
	assert.NoError(t, err)

	t.Cleanup(func() {
		assert.NoError(t, testcontainers.TerminateContainer(ctr))
	})

	return ctr
}
