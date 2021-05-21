package api

import (
	"context"
	"fmt"
	"grpc-go-templete/internal/user/helper"
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
	userService *service.UserService
}

func newUserGrpcServer(userService *service.UserService) *userGrpcServer {
	return &userGrpcServer{
		userService: userService,
	}
}

func InitUserGrpcServer(host, port string, userService *service.UserService) (conn *grpc.ClientConn) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalln("User gRPC - Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the service to the server
	user_pb.RegisterUsersServer(s, newUserGrpcServer(userService))
	// Serve gRPC Server
	log.Printf("User gRPC - Started on %s:%s", host, port)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	conn, err = grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%s", host, port),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("User gRPC - Failed to dial server:", err)
	}
	return
}

func InitGrpcGetway(host, port string, conn *grpc.ClientConn) (gwServer *http.Server) {
	// Create http server
	gwmux := runtime.NewServeMux()
	// Attach the server dto server
	err := user_pb.RegisterUsersHandler(context.Background(), gwmux, conn)
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

func (s *userGrpcServer) CreateUser(ctx context.Context, in *user_pb.User) (*user_pb.User, error) {
	user, err := s.userService.CreateUser(helper.ToUser(in))
	if err != nil {
		return nil, err
	}
	resp := helper.ToUserGRPC(user)
	return resp, nil
}

func (s *userGrpcServer) GetUser(ctx context.Context, in *user_pb.UserRequest) (*user_pb.UserResponse, error) {
	resp := &user_pb.UserResponse{}
	return resp, nil
}

func (s *userGrpcServer) GetUserList(ctx context.Context, in *user_pb.UserListRequest) (*user_pb.UserListResponse, error) {
	resp := &user_pb.UserListResponse{}
	return resp, nil
}
