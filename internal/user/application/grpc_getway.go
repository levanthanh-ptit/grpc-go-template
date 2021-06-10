package application

import (
	"context"
	"fmt"
	"grpc-go-templete/pkg/pb/user_pb"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// GrpcGetwayServer is getway server
type GrpcGetwayServer struct {
	host   string
	port   string
	gwmux  *runtime.ServeMux
	server *http.Server

	targetAddr string
	targetConn *grpc.ClientConn
}

// NewGrpcGetway constructor
func NewGrpcGetway(targetAddr, host, port string) *GrpcGetwayServer {
	gwmux := runtime.NewServeMux()
	return &GrpcGetwayServer{
		targetAddr: targetAddr,
		host:       host,
		port:       port,
		gwmux:      gwmux,
	}
}

// Serve run server method
func (s GrpcGetwayServer) Serve() {
	var err error = nil
	s.targetConn, err = grpc.DialContext(
		context.Background(),
		s.targetAddr,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("User gRPC - Failed to dial server:", err)
	}
	err = user_pb.RegisterUsersHandler(context.Background(), s.gwmux, s.targetConn)
	if err != nil {
		log.Fatalln("User gRPC-Gateway - Failed to register gateway:", err)
	}
	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.host, s.port),
		Handler: s.gwmux,
	}
	log.Printf("User gRPC-Gateway - Started on http://%s:%s", s.host, s.port)
	log.Fatalln(s.server.ListenAndServe())
}

// Close stop server method
func (s GrpcGetwayServer) Close() {

}
