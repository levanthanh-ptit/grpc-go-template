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

type userGrpcServer struct {
	user_pb.UnimplementedUserServer
}

func newUserGrpcServer() *userGrpcServer {
	return &userGrpcServer{}
}

func InitUserGrpcServer() (conn *grpc.ClientConn) {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("User gRPC - Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the service to the server
	user_pb.RegisterUserServer(s, newUserGrpcServer())
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
	err := user_pb.RegisterUserHandler(context.Background(), gwmux, conn)
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

func (s *userGrpcServer) GetUser(ctx context.Context, in *user_pb.UserRequest) (*user_pb.UserResponse, error) {
	return &user_pb.UserResponse{Id: in.Id, Name: "aaaaaaa"}, nil
}
