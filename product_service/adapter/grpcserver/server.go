package grpcserver

import (
	pb "github.com/meokg456/ecommerce/proto/product"

	"github.com/meokg456/productservice/domain/product"
	"github.com/meokg456/productservice/pkg/config"
	"go.uber.org/zap"
)

type Server struct {
	pb.UnimplementedProductServiceServer

	Config config.Config
	Logger *zap.SugaredLogger

	ProductStore product.Storage
}

func New(config *config.Config) *Server {
	return &Server{
		Config: *config,
	}
}
