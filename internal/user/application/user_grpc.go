package application

import (
	"context"
	"fmt"
	"grpc-go-templete/internal/user/service"
	"grpc-go-templete/pkg/pb/product_pb"
	"grpc-go-templete/pkg/pb/user_pb"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type usersGrpcServer struct {
	user_pb.UnimplementedUsersServer
	host string
	port string

	// Clients
	ProductsClient product_pb.ProductsClient

	// Services
	userService *service.UserService
}

func NewUserGrpcServer(host, port string, userService *service.UserService) *usersGrpcServer {
	return &usersGrpcServer{
		host: host,
		port: port,

		userService: userService,
	}
}

func (usersServer *usersGrpcServer) StartUserGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", usersServer.host, usersServer.port))
	if err != nil {
		log.Fatalln("User gRPC - Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the service to the server
	user_pb.RegisterUsersServer(s, usersServer)
	// Serve gRPC Server
	log.Printf("User gRPC - Started on %s:%s", usersServer.host, usersServer.port)
	log.Fatalln(s.Serve(lis))
}

func (usersServer *usersGrpcServer) StartGrpcGetwayServer(host, port string) (gwServer *http.Server) {
	// Dial to GRPC server
	grpcConn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%s", usersServer.host, usersServer.port),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("User gRPC - Failed to dial server:", err)
	}
	defer grpcConn.Close()
	// Create http server
	gwmux := runtime.NewServeMux()
	// Attach the server dto server
	err = user_pb.RegisterUsersHandler(context.Background(), gwmux, grpcConn)
	if err != nil {
		log.Fatalln("User gRPC-Gateway - Failed to register gateway:", err)
	}
	gwServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: gwmux,
	}
	log.Printf("User gRPC-Gateway - Started on http://%s:%s", host, port)
	log.Fatalln(gwServer.ListenAndServe())
	return
}
