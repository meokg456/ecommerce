package grpcserver

import (
	"context"

	pb "proto/product"

	"github.com/meokg456/productservice/domain/product"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *Server) AddProduct(ctx context.Context, request *pb.Product) (*pb.Product, error) {

	p := product.Product{
		Title:        request.Title,
		Descriptions: request.Descriptions,
		Category:     request.Category,
		Images:       request.Images,
		AdditionInfo: request.AdditionInfo.AsMap(),
	}

	err := s.ProductStore.AddProduct(&p)

	if err != nil {
		s.Logger.Errorf("add product grpc: failed to add product %v", request)
		return nil, err
	}

	return &pb.Product{
		Id:           p.Id,
		Title:        p.Title,
		Descriptions: p.Descriptions,
		Category:     p.Category,
		Images:       p.Images,
		AdditionInfo: &structpb.Struct{},
	}, nil
}
