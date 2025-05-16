package grpcservice

import (
	"context"

	pbcommon "github.com/meokg456/ecommerce/proto/common"
	pb "github.com/meokg456/ecommerce/proto/product"

	"github.com/meokg456/productmanagement/domain/common"
	"github.com/meokg456/productmanagement/domain/product"
	"google.golang.org/protobuf/types/known/structpb"
)

type ProductService struct {
	client pb.ProductServiceClient
}

func NewProductService(productClient pb.ProductServiceClient) *ProductService {
	return &ProductService{
		client: productClient,
	}
}

func (p ProductService) AddProduct(product *product.Product) error {
	additionInfo, err := structpb.NewStruct(product.AdditionInfo)
	if err != nil {
		return err
	}

	result, err := p.client.AddProduct(context.Background(), &pb.Product{
		Title:        product.Title,
		Descriptions: product.Descriptions,
		Category:     product.Category,
		Images:       product.Images,
		AdditionInfo: additionInfo,
		MerchantId:   int64(product.MerchantId),
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

	_, err = p.client.UpdateProduct(context.Background(), &pb.Product{
		Id:           product.Id,
		Title:        product.Title,
		Descriptions: product.Descriptions,
		Category:     product.Category,
		Images:       product.Images,
		AdditionInfo: additionInfo,
		MerchantId:   int64(product.MerchantId),
	})

	if err != nil {
		return err
	}

	return nil
}

func (p ProductService) DeleteProduct(merchantId int, id string) error {
	_, err := p.client.DeleteProduct(context.Background(), &pb.DeleteProductRequest{
		Id:         id,
		MerchantId: int64(merchantId),
	})

	if err != nil {
		return err
	}

	return nil
}

func (p ProductService) GetProductsByMerchantId(merchantId int, page common.Page) ([]product.Product, string, error) {
	var products []product.Product

	response, err := p.client.GetProductsByMerchantId(context.Background(), &pb.GetProductsByMerchantIdRequest{
		MerchantId: int64(merchantId),
		Page: &pbcommon.Page{
			LastKeyOffset: page.LastKeyOffset,
			Limit:         int32(page.Limit),
		},
	})

	if err != nil {
		return nil, "", err
	}

	for _, p := range response.Products {
		products = append(products, product.NewProductWithId(
			p.Id,
			p.Title,
			p.Descriptions,
			p.Category,
			p.Images,
			p.AdditionInfo.AsMap(),
			int(p.MerchantId),
		))
	}

	return products, response.LastKey, nil
}
