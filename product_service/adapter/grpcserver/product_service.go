package grpcserver

import (
	"context"

	pb "github.com/meokg456/ecommerce/proto/product"

	"github.com/meokg456/productservice/domain/common"
	"github.com/meokg456/productservice/domain/product"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *Server) GetProductsByMerchantId(ctx context.Context, request *pb.GetProductsByMerchantIdRequest) (*pb.GetProductsByMerchantIdResponse, error) {

	products, lastKey, err := s.ProductStore.GetProductsByMerchantId(int(request.MerchantId), common.Page{
		Page:          int(request.Page.Page),
		LastKeyOffset: request.Page.LastKeyOffset,
		Limit:         int(request.Page.Limit),
	})

	if err != nil {
		s.Logger.Errorf("get products by merchant id grpc: failed to get products %v: %v", request, err)
		return nil, err
	}

	var result []*pb.Product

	for _, p := range products {
		result = append(result, &pb.Product{
			Id:           p.Id,
			Title:        p.Title,
			Descriptions: p.Descriptions,
			Category:     p.Category,
			Images:       p.Images,
			AdditionInfo: &structpb.Struct{},
			MerchantId:   int64(p.MerchantId),
		})
	}

	return &pb.GetProductsByMerchantIdResponse{
		Products: result,
		LastKey:  lastKey,
	}, nil
}

func (s *Server) AddProduct(ctx context.Context, request *pb.Product) (*pb.Product, error) {

	p := product.Product{
		Title:        request.Title,
		Descriptions: request.Descriptions,
		Category:     request.Category,
		Images:       request.Images,
		AdditionInfo: request.AdditionInfo.AsMap(),
		MerchantId:   int(request.MerchantId),
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
		MerchantId:   int64(p.MerchantId),
	}, nil
}

func (s *Server) UpdateProduct(ctx context.Context, request *pb.Product) (*pb.Product, error) {

	p := product.Product{
		Id:           request.Id,
		Title:        request.Title,
		Descriptions: request.Descriptions,
		Category:     request.Category,
		Images:       request.Images,
		AdditionInfo: request.AdditionInfo.AsMap(),
		MerchantId:   int(request.MerchantId),
	}

	err := s.ProductStore.UpdateProduct(&p)

	if err != nil {
		s.Logger.Errorf("update product grpc: failed to update product %v", request)
		return nil, err
	}

	return &pb.Product{
		Id:           p.Id,
		Title:        p.Title,
		Descriptions: p.Descriptions,
		Category:     p.Category,
		Images:       p.Images,
		AdditionInfo: &structpb.Struct{},
		MerchantId:   int64(p.MerchantId),
	}, nil
}

func (s *Server) DeleteProduct(ctx context.Context, request *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {

	err := s.ProductStore.DeleteProduct(int(request.MerchantId), request.Id)

	if err != nil {
		s.Logger.Errorf("delete product grpc: failed to delete product %v", request)
		return nil, err
	}

	return &pb.DeleteProductResponse{
		Id: request.Id,
	}, nil
}
