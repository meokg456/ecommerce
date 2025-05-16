package grpcservice

import (
	"context"

	pb "github.com/meokg456/ecommerce/proto/inventory"
	"github.com/meokg456/productmanagement/domain/inventory"
)

type InventoryStore struct {
	client pb.InventoryServiceClient
}

func NewInventoryService(client pb.InventoryServiceClient) *InventoryStore {
	return &InventoryStore{
		client: client,
	}
}

func (i *InventoryStore) SaveInventory(inv *inventory.Inventory) error {
	updatedInventory, err := i.client.SaveInventory(context.Background(), &pb.Inventory{
		ProductId: inv.ProductId,
		Types:     inv.Types,
		Quantity:  int64(inv.Quantity),
	})

	if err != nil {
		return err
	}

	inv.Quantity = int(updatedInventory.Quantity)

	return nil
}

func (i *InventoryStore) GetInventory(productId string, t []string) (*inventory.Inventory, error) {
	result, err := i.client.GetInventory(context.Background(), &pb.GetInventoryRequest{
		ProductId: productId,
		Types:     t,
	})

	if err != nil {
		return nil, err
	}

	inv := inventory.NewInventory(result.ProductId, result.Types, int(result.Quantity))

	return &inv, nil
}
