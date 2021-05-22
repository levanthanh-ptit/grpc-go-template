package application

import (
	"context"
	"fmt"
	"grpc-go-templete/internal/product/service"
	"grpc-go-templete/pkg/pb/product_pb"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type productsGrpcServer struct {
	product_pb.UnimplementedProductsServer
	host string
	port string

	productService *service.ProductService
}

func NewProductsGrpcServer(host, port string, productService *service.ProductService) *productsGrpcServer {
	return &productsGrpcServer{
		host: host,
		port: port,

		productService: productService,
	}
}

func (productsServer *productsGrpcServer) StartGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", productsServer.host, productsServer.port))
	if err != nil {
		log.Fatalln("Product gRPC - Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the service to the server
	product_pb.RegisterProductsServer(s, productsServer)
	// Serve gRPC Server
	log.Printf("Product gRPC - Started on %s:%s", productsServer.host, productsServer.port)
	log.Fatalln(s.Serve(lis))
}

func (productsServer *productsGrpcServer) StartGrpcGetway(host, port string) (gwServer *http.Server) {
	// Dial to GRPC server
	grpcConn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%s", productsServer.host, productsServer.port),
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
		Handler: gwmux,
	}
	log.Printf("Product gRPC-Gateway - Started on http://%s:%s", host, port)
	log.Fatalln(gwServer.ListenAndServe())
	return
}
