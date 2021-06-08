package main

import (
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
		panic("DB not connect")
	}

	// Init repositories
	productRepo := infrastructure.NewProductPersistance(mongoProvider.GetDatabase("product_dev"))

	// Init services
	productService := service.NewProductService(productRepo)

	// Application enums
	host := "localhost"
	grpcPort := "8082"
	grpcGwPort := "8092"
	userServerAddress := "localhost:8091"

	// Init Application
	productGrpcServer := application.NewProductsGrpcServer(host, grpcPort, productService)

	// Start GRPC server
	go productGrpcServer.StartGrpcServer()

	// Init Clients
	userConn, err := productGrpcServer.RegisterUsersClient(userServerAddress)
	if err != nil {
		log.Fatalln(err)
	}

	var closeClients = func() {
		userConn.Close()
	}
	defer closeClients()

	// Start GRPC Getway
	productGrpcServer.StartGrpcGetway(host, grpcGwPort)

}
