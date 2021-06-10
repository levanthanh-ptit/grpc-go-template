package application

import (
	"fmt"
	"grpc-go-templete/internal/user/service"
	"grpc-go-templete/pkg/pb/product_pb"
	"grpc-go-templete/pkg/pb/user_pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	user_pb.UnimplementedUsersServer
	name string
	host string
	port string

	server *grpc.Server

	// Clients
	ProductsClient product_pb.ProductsClient

	// Services
	userService *service.UserService
	authService *service.AuthService
}

func NewGrpcServer(host, port string, userService *service.UserService, authService *service.AuthService) *GrpcServer {
	return &GrpcServer{
		host: host,
		port: port,

		userService: userService,
		authService: authService,
	}
}

func (s *GrpcServer) Serve() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		log.Fatalln("User gRPC - Failed to listen:", err)
	}
	// Create a gRPC server object
	s.server = grpc.NewServer()
	// Attach the service to the server
	user_pb.RegisterUsersServer(s.server, s)
	// Serve gRPC Server
	log.Printf("User gRPC - Started on %s:%s", s.host, s.port)
	log.Fatalln(s.Serve(lis))
}
