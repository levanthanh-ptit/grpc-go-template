package application

import (
	"grpc-go-templete/internal/product/service"
	"grpc-go-templete/pkg/pb/product_pb"

	"github.com/levanthanh-ptit/go-grpc-kit/server"
	"google.golang.org/grpc"
)

// GrpcServer server object
type GrpcServer struct {
	server.GrpcServer
	product_pb.UnimplementedProductsServer

	// Services
	productService *service.ProductService
}

// NewGrpcServer constructor
func NewGrpcServer(
	host,
	port string,
	productService *service.ProductService,
) *GrpcServer {
	base := server.NewGrpcServer("user_grpc_server")
	s := &GrpcServer{
		GrpcServer:     *base,
		productService: productService,
	}
	s.WithHost(host)
	s.WithPort(port)
	s.WithGprpcRegister(s.RegisterServer)
	return s
}

// RegisterServer register server
func (s *GrpcServer) RegisterServer(server *grpc.Server) (err error) {
	product_pb.RegisterProductsServer(server, s)
	return
}

func (s *GrpcServer) Close() {
}
