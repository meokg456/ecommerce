package grpcserver

import (
	pb "github.com/meokg456/ecommerce/proto/inventory"

	"github.com/meokg456/inventoryservice/domain/inventory"
	"github.com/meokg456/inventoryservice/pkg/config"
	"go.uber.org/zap"
)

type Server struct {
	pb.UnimplementedInventoryServiceServer

	Config config.Config
	Logger *zap.SugaredLogger

	InventoryStore inventory.Storage
}

func New(config *config.Config) *Server {
	return &Server{
		Config: *config,
	}
}
