package grpcservice

import (
	pb "proto/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProductServiceClient(host string) (pb.ProductServiceClient, error) {
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pb.NewProductServiceClient(conn), nil
}
