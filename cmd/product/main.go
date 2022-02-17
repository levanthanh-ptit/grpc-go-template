package main

import (
	"fmt"
	"grpc-go-templete/internal/product/application"
	"grpc-go-templete/internal/product/infrastructure"
	"grpc-go-templete/internal/product/service"
	"log"

	"github.com/levanthanh-ptit/go-ez/ez_provider"
)

func main() {
	// Init Database
	mongoProvider, err := ez_provider.NewMongoProvider("mongodb://localhost:27017")
	if err != nil {
		log.Fatalln(err)
	}

	// Init repositories
	productRepo := infrastructure.NewProductRepositoryMongo(mongoProvider.GetDatabase("product_dev"))

	// Init services
	productService := service.NewProductService(productRepo)

	// Application enums
	host := "localhost"
	grpcPort := "8082"
	grpcGwPort := "8092"
	userServerAddress := "localhost:8081"

	// Init Application
	grpcServer := application.NewGrpcServer(host, grpcPort, productService)
	grpcGateway := application.NewGrpcGateway(fmt.Sprintf("%s:%s", host, grpcPort), host, grpcGwPort)
	grpcGateway.RegisterUsersClient(userServerAddress)

	// Start GRPC server
	go grpcServer.ServeTCP()

	// Start GRPC Gateway
	defer grpcGateway.Close()
	grpcGateway.Serve()
}
