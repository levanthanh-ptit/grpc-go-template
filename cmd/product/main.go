package main

import (
	"grpc-go-templete/internal/product/application"
	"grpc-go-templete/internal/product/infrastructure"
	"grpc-go-templete/internal/product/service"
	"grpc-go-templete/pkg/provider"
)

func main() {
	// Init Database
	mongoProvider, err := provider.NewMongoProvider("mongodb://localhost:27017")
	if err != nil {
		panic("DB not connect")
	}

	// Init repositories
	productRepo := infrastructure.NewProductPersistance(mongoProvider.GetDatabase("product_dev"))

	// Init services
	productService := service.NewProductService(productRepo)

	// Application enums
	host := "localhost"
	grpcPort := "8083"
	grpcGwPort := "8093"

	// Init Application
	productGrpcServer := application.NewProductsGrpcServer(host, grpcPort, productService)

	// Start GRPC server
	go productGrpcServer.StartGrpcServer()

	// Start GRPC Getway
	productGrpcServer.StartGrpcGetway(host, grpcGwPort)

}
