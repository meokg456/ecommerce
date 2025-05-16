package grpcserver

import (
	"context"
	"log"
	"net"
	"runtime/debug"

	pb "github.com/meokg456/ecommerce/proto/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/meokg456/productservice/domain/product"
	"github.com/meokg456/productservice/pkg/config"
	"go.uber.org/zap"
)

type Server struct {
	pb.UnimplementedProductServiceServer

	grpcServer *grpc.Server

	Config config.Config
	Logger *zap.SugaredLogger

	ProductStore product.Storage
}

func New(config *config.Config) *Server {
	server := &Server{
		Config: *config,
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(PanicRecoveryInterceptor),
	)

	server.grpcServer = grpcServer

	pb.RegisterProductServiceServer(grpcServer, server)

	return server
}

func (s *Server) Serve(lis net.Listener) error {
	return s.grpcServer.Serve(lis)
}

// PanicRecoveryInterceptor is a unary interceptor for recovering from panics
func PanicRecoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			// Log the panic and stack trace
			log.Printf("Panic occurred: %v\n", r)
			log.Printf("Stack trace: %s\n", debug.Stack())

			// Return a gRPC error to the client
			err = status.Errorf(codes.Internal, "Internal server error")
		}
	}()

	// Call the handler to proceed with the normal execution of the RPC
	return handler(ctx, req)
}
