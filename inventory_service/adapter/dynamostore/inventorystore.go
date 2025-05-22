package dynamostore

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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

func (i *InventoryStore) SaveInventory(inv *inventory.Inventory) error {
	key := inventory.HashInventory(inv.ProductId, inv.Types)

	output, err := i.db.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(dbconst.InventoryTableName),
		Key: map[string]types.AttributeValue{
			dbconst.InventoryPK: &types.AttributeValueMemberS{Value: key},
		},
		UpdateExpression:    aws.String("ADD Quantity :q"),
		ConditionExpression: aws.String("(attribute_not_exists(Quantity) AND :q >= :zero) OR Quantity >= :min"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":q":    &types.AttributeValueMemberN{Value: strconv.Itoa(inv.Quantity)},
			":min":  &types.AttributeValueMemberN{Value: strconv.Itoa(-inv.Quantity)},
			":zero": &types.AttributeValueMemberN{Value: strconv.Itoa(0)},
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	})

	if err != nil {
		return err
	}

	quantity, err := strconv.Atoi(output.Attributes[dbconst.InventoryQuantity].(*types.AttributeValueMemberN).Value)
	if err != nil {
		return err
	}

	inv.Quantity = quantity

	return nil
}

func (i *InventoryStore) GetInventory(productId string, t []string) (*inventory.Inventory, error) {
	key := inventory.HashInventory(productId, t)

	output, err := i.db.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(dbconst.InventoryTableName),
		Key: map[string]types.AttributeValue{
			dbconst.InventoryPK: &types.AttributeValueMemberS{Value: key},
		},
	})

	if err != nil {
		return nil, err
	}

	var data InventoryData

	err = attributevalue.UnmarshalMap(output.Item, &data)
	if err != nil {
		return nil, err
	}

	inv := inventory.NewInventory(productId, t, data.Quantity)

	return &inv, nil
}
