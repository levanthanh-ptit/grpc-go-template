package application

import (
	"context"
	"fmt"
	"grpc-go-templete/internal/user/service"
	"grpc-go-templete/pkg/pb/user_pb"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type userGrpcServer struct {
	user_pb.UnimplementedUsersServer
	conn *grpc.ClientConn

	userService *service.UserService
}

func NewUserGrpcServer(userService *service.UserService) *userGrpcServer {
	return &userGrpcServer{
		userService: userService,
	}
}

func (userServer *userGrpcServer) StartUserGrpcServer(host, port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalln("User gRPC - Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the service to the server
	user_pb.RegisterUsersServer(s, userServer)
	// Serve gRPC Server
	log.Printf("User gRPC - Started on %s:%s", host, port)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	userServer.conn, err = grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%s", host, port),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("User gRPC - Failed to dial server:", err)
	}
}

func (userServer *userGrpcServer) StartGrpcGetwayServer(host, port string) (gwServer *http.Server) {
	// Create http server
	gwmux := runtime.NewServeMux()
	// Attach the server dto server
	err := user_pb.RegisterUsersHandler(context.Background(), gwmux, userServer.conn)
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
