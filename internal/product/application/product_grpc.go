package application

import (
	"context"
	"fmt"
	"grpc-go-templete/internal/product/service"
	"grpc-go-templete/pkg/pb/product_pb"
	"grpc-go-templete/pkg/pb/user_pb"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type ProductsGrpcServer struct {
	product_pb.UnimplementedProductsServer
	host string
	port string

	// Clients
	UsersClient user_pb.UsersClient

	// Services
	productService *service.ProductService
}

func NewProductsGrpcServer(host, port string, productService *service.ProductService) *ProductsGrpcServer {
	return &ProductsGrpcServer{
		host: host,
		port: port,

		productService: productService,
	}
}

func (s *ProductsGrpcServer) StartGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		log.Fatalln("Product gRPC - Failed to listen:", err)
	}
	// Create a gRPC server object
	server := grpc.NewServer()
	// Attach the service to the server
	product_pb.RegisterProductsServer(server, s)
	// Serve gRPC Server
	log.Printf("Product gRPC - Started on %s:%s", s.host, s.port)
	log.Fatalln(server.Serve(lis))
}

func (s *ProductsGrpcServer) StartGrpcGetway(host, port string) (gwServer *http.Server) {
	// Dial to GRPC server
	grpcConn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%s", s.host, s.port),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Product gRPC - Failed to dial server:", err)
	}
	defer grpcConn.Close()
	// Create http server
	gwmux := runtime.NewServeMux()
	// Attach the server dto server
	err = product_pb.RegisterProductsHandler(context.Background(), gwmux, grpcConn)
	if err != nil {
		log.Fatalln("Product gRPC-Gateway - Failed to register gateway:", err)
	}
	gwServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", host, port),
		Handler: s.AuthGuard(gwmux),
	}
	log.Printf("Product gRPC-Gateway - Started on http://%s:%s", host, port)
	log.Fatalln(gwServer.ListenAndServe())
	return
}
