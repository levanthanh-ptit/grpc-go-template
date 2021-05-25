package application

import (
	"fmt"
	"grpc-go-templete/internal/auth/service"
	"grpc-go-templete/pkg/pb/auth_pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type authGrpcServer struct {
	auth_pb.UnimplementedAuthServer
	host string
	port string

	authService *service.AuthService
}

func NewAuthGrpcServer(host, port string, authService *service.AuthService) *authGrpcServer {
	return &authGrpcServer{
		host: host,
		port: port,

		authService: authService,
	}
}

func (authServer *authGrpcServer) StartGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", authServer.host, authServer.port))
	if err != nil {
		log.Fatalln("Auth gRPC - Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the service to the server
	auth_pb.RegisterAuthServer(s, authServer)
	// Serve gRPC Server
	log.Printf("Auth gRPC - Started on %s:%s", authServer.host, authServer.port)
	log.Fatalln(s.Serve(lis))
}
