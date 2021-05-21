package main

import (
	"grpc-go-templete/internal/product/api"
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

	host := "localhost"
	grpcPort := "8082"
	grpcGwPort := "8092"

	// Init Application
	conn := api.InitGrpcServer(host, grpcPort, productService)
	api.InitGrpcGetway(host, grpcGwPort, conn)
}
