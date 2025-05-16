package grpcserver

import (
	"context"

	pb "github.com/meokg456/ecommerce/proto/inventory"
	"github.com/meokg456/inventoryservice/domain/inventory"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s Server) SaveInventory(context context.Context, inv *pb.Inventory) (*emptypb.Empty, error) {
	err := s.InventoryStore.SaveInventory(inventory.NewInventory(
		inv.ProductId,
		inv.Types,
		int(inv.Quantity),
	))

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
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
