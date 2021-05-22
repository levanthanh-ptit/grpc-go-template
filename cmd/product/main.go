package main

import (
	"grpc-go-templete/internal/product/application"
	"grpc-go-templete/internal/product/persistance"
	"grpc-go-templete/internal/product/service"
)

func main() {
	// Init Database
	mongoProvider, err := persistance.NewMongoProvider("mongodb://localhost:27017")
	if err != nil {
		panic("DB not connect")
	}

	// Init repositories
	productRepo := persistance.NewProductPersistance(mongoProvider.GetDatabase("product_dev"))

	// Init services
	productService := service.NewProductService(productRepo)

	// Application enums
	host := "localhost"
	grpcPort := "8082"
	grpcGwPort := "8092"

	// Init Application
	productGrpcServer := application.NewProductsGrpcServer(host, grpcPort, productService)

	// Start GRPC server
	go productGrpcServer.StartGrpcServer()

	// Start GRPC Getway
	productGrpcServer.StartGrpcGetway(host, grpcGwPort)

}
