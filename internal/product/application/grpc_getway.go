package application

import (
	"context"
	"grpc-go-templete/pkg/pb/product_pb"
	"grpc-go-templete/pkg/pb/user_pb"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/levanthanh-ptit/go-grpc-kit/server"
	"google.golang.org/grpc"
)

// GrpcGetwayServer is getway server
type GrpcGetwayServer struct {
	server.GrpcGetwayServer

	targetAddr string
	targetConn *grpc.ClientConn

	usersClient user_pb.UsersClient
	usersConn   *grpc.ClientConn
}

// NewGrpcGetway constructor
func NewGrpcGetway(targetAddr, host, port string) *GrpcGetwayServer {
	base := server.NewGrpcGetwayServer("user_grpc_getway")
	s := &GrpcGetwayServer{
		GrpcGetwayServer: *base,
		targetAddr:       targetAddr,
	}
	s.WithHost(host)
	s.WithPort(port)
	s.WithClientRegister(s.RegisterGrpcClient)
	s.WithHandler(s.AuthGuard)
	return s
}

// RegisterGrpcClient attach gRPC
func (s *GrpcGetwayServer) RegisterGrpcClient(gwmux *runtime.ServeMux) {
	var err error = nil
	s.targetConn, err = grpc.DialContext(
		context.Background(),
		s.targetAddr,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Product gRPC - Failed to dial server:", err)
	}
	err = product_pb.RegisterProductsHandler(context.Background(), gwmux, s.targetConn)
	if err != nil {
		log.Fatalln("Product gRPC-Gateway - Failed to register gateway:", err)
	}
}

// Close close clients...
func (s *GrpcGetwayServer) Close() {
	s.targetConn.Close()
	s.usersConn.Close()
}
