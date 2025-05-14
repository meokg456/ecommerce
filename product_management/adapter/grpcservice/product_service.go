package grpcservice

import (
	"context"
	pb "proto/product"

	"github.com/meokg456/productmanagement/domain/product"
	"google.golang.org/protobuf/types/known/structpb"
)

type ProductService struct {
	productClient pb.ProductServiceClient
}

func NewProductService(productClient pb.ProductServiceClient) *ProductService {
	return &ProductService{
		productClient: productClient,
	}
}

func (p ProductService) AddProduct(product *product.Product) error {
	additionInfo, err := structpb.NewStruct(product.AdditionInfo)
	if err != nil {
		return err
	}

	result, err := p.productClient.AddProduct(context.Background(), &pb.Product{
		Title:        product.Title,
		Descriptions: product.Descriptions,
		Category:     product.Category,
		Images:       product.Images,
		AdditionInfo: additionInfo,
	})

	if err != nil {
		return err
	}
	product.Id = result.Id

	return nil
}
