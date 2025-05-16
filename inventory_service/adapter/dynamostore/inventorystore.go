package dynamostore

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/meokg456/inventoryservice/dbconst"
	"github.com/meokg456/inventoryservice/domain/inventory"
)

type InventoryStore struct {
	db *dynamodb.Client
}

func NewInventoryStore(db *dynamodb.Client) *InventoryStore {
	return &InventoryStore{
		db: db,
	}
}

func (i *InventoryStore) SaveInventory(inv inventory.Inventory) error {
	key := inventory.HashInventory(inv)

	_, err := i.db.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(dbconst.InventoryTableName),
		Item: map[string]types.AttributeValue{
			dbconst.InventoryPK:       &types.AttributeValueMemberS{Value: key},
			dbconst.InventoryQuantity: &types.AttributeValueMemberN{Value: strconv.Itoa(inv.Quantity)},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
