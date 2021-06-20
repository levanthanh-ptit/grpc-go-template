package application

import (
	"context"
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
	return s
}

// RegisterGrpcClient attach gRPC
func (s *GrpcGetwayServer) RegisterGrpcClient(gwmux *runtime.ServeMux) (err error) {
	s.targetConn, err = grpc.DialContext(
		context.Background(),
		s.targetAddr,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Println("User gRPC - Failed to dial server")
		return
	}
	err = user_pb.RegisterUsersHandler(context.Background(), gwmux, s.targetConn)
	if err != nil {
		log.Println("User gRPC-Gateway - Failed to register gateway")
		return
	}

	return
}

// Close close clients...
func (s *GrpcGetwayServer) Close() {
	s.targetConn.Close()
}
