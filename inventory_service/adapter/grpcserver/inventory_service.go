package grpcserver

import (
	"context"

	pb "github.com/meokg456/ecommerce/proto/inventory"
	"github.com/meokg456/inventoryservice/domain/inventory"
)

func (s Server) SaveInventory(context context.Context, request *pb.Inventory) (*pb.Inventory, error) {
	inv := inventory.NewInventory(
		request.ProductId,
		request.Types,
		int(request.Quantity),
	)
	err := s.InventoryStore.SaveInventory(&inv)

	if err != nil {
		return nil, err
	}

	return &pb.Inventory{
		ProductId: inv.ProductId,
		Types:     inv.Types,
		Quantity:  int64(inv.Quantity),
	}, nil
}

func (s Server) GetInventory(context context.Context, req *pb.GetInventoryRequest) (*pb.Inventory, error) {
	inv, err := s.InventoryStore.GetInventory(req.ProductId, req.Types)

	if err != nil {
		return nil, err
	}

	return &pb.Inventory{
		ProductId: inv.ProductId,
		Types:     inv.Types,
		Quantity:  int64(inv.Quantity),
	}, nil
}
