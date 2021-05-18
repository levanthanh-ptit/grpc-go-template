package api

import (
	"context"
	"grpc-go-templete/pkg/pb/user_pb"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type usersGrpcServer struct {
	user_pb.UnimplementedUsersServer
}

func newUsersGrpcServer() *usersGrpcServer {
	return &usersGrpcServer{}
}

func InitUserGrpcServer() (conn *grpc.ClientConn) {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("User gRPC - Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the service to the server
	user_pb.RegisterUsersServer(s, newUsersGrpcServer())
	// Serve gRPC Server
	log.Println("User gRPC - Started on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	conn, err = grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("User gRPC - Failed to dial server:", err)
	}
	return
}

func InitGrpcGetway(conn *grpc.ClientConn) (gwServer *http.Server) {
	// Create http server
	gwmux := runtime.NewServeMux()
	// Attach the server dto server
	err := user_pb.RegisterUsersHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("User gRPC-Gateway - Failed to register gateway:", err)
	}
	gwServer = &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	log.Println("User gRPC-Gateway - Started on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
	return
}

func (s *usersGrpcServer) GetUser(ctx context.Context, in *user_pb.UserRequest) (*user_pb.UserResponse, error) {
	return &user_pb.UserResponse{Data: &user_pb.User{}}, nil
}
