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

// GrpcGatewayServer is getway server
type GrpcGatewayServer struct {
	server.GrpcGatewayServer

	targetAddr string
	targetConn *grpc.ClientConn

	usersClient user_pb.UsersClient
	usersConn   *grpc.ClientConn
}

// NewGrpcGateway constructor
func NewGrpcGateway(targetAddr, host, port string) *GrpcGatewayServer {
	base := server.NewGrpcGatewayServer("user_grpc_getway")
	s := &GrpcGatewayServer{
		GrpcGatewayServer: *base,
		targetAddr:        targetAddr,
	}
	s.WithHost(host)
	s.WithPort(port)
	s.WithClientRegister(s.RegisterGrpcClient)
	s.WithHTTPHandler(s.AuthGuard)
	return s
}

// RegisterGrpcClient attach gRPC
func (s *GrpcGatewayServer) RegisterGrpcClient(gwmux *runtime.ServeMux) (err error) {
	s.targetConn, err = grpc.DialContext(
		context.Background(),
		s.targetAddr,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Println("Product gRPC - Failed to dial server")
		return
	}
	err = product_pb.RegisterProductsHandler(context.Background(), gwmux, s.targetConn)
	if err != nil {
		log.Println("Product gRPC-Gateway - Failed to register gateway")
		return
	}

	return
}

// Close close clients...
func (s *GrpcGatewayServer) Close() {
	s.targetConn.Close()
	s.usersConn.Close()
}
