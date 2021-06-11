package application

import (
	"grpc-go-templete/internal/user/service"
	"grpc-go-templete/pkg/pb/product_pb"
	"grpc-go-templete/pkg/pb/user_pb"

	"github.com/levanthanh-ptit/go-grpc-kit/server"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	server.GrpcServer
	user_pb.UnimplementedUsersServer

	// Clients
	ProductsClient product_pb.ProductsClient

	// Services
	userService *service.UserService
	authService *service.AuthService
}

// NewGrpcServer constructor
func NewGrpcServer(
	host,
	port string,
	userService *service.UserService,
	authService *service.AuthService,
) *GrpcServer {
	base := server.NewGrpcServer("user_grpc_server")
	s := &GrpcServer{
		GrpcServer:  *base,
		userService: userService,
		authService: authService,
	}
	s.WithHost(host).WithPort(port).WithGprpcRegister(s.RegisterServer)
	return s
}

// RegisterServer register func
func (s *GrpcServer) RegisterServer(server *grpc.Server) {
	user_pb.RegisterUsersServer(server, s)
}
