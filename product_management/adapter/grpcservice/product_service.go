package grpcservice

import (
	"context"

	pb "github.com/meokg456/ecommerce/proto/product"

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

func (p ProductService) UpdateProduct(product *product.Product) error {
	additionInfo, err := structpb.NewStruct(product.AdditionInfo)
	if err != nil {
		return err
	}

	_, err = p.productClient.UpdateProduct(context.Background(), &pb.Product{
		Id:           product.Id,
		Title:        product.Title,
		Descriptions: product.Descriptions,
		Category:     product.Category,
		Images:       product.Images,
		AdditionInfo: additionInfo,
	})

	if err != nil {
		return err
	}

	return nil
}

func (p ProductService) DeleteProduct(id string) error {
	_, err := p.productClient.DeleteProduct(context.Background(), &pb.DeleteProductRequest{
		Id: id,
	})

	if err != nil {
		return err
	}

	return nil
}
