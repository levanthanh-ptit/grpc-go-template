package main

import (
	"grpc-go-templete/internal/user/application"
	"grpc-go-templete/internal/user/persistance"
	"grpc-go-templete/internal/user/service"
	"log"
)

func main() {
	// Init Database
	mongoProvider, err := persistance.NewMongoProvider("mongodb://localhost:27017")
	if err != nil {
		panic("DB not connect")
	}

	// Init repositories
	userRepo := persistance.NewUserPersistance(mongoProvider.GetDatabase("user_dev"))

	// Init services
	userService := service.NewUserService(userRepo)

	// Application enums
	host := "localhost"
	grpcPort := "8081"
	grpcGwPort := "8091"
	productServerAddress := "localhost:8082"

	// Init Application
	userGrpcServer := application.NewUserGrpcServer(userService)

	// Init Clients
	productConn, err := userGrpcServer.RegisterProductsClient(productServerAddress)
	if err != nil {
		log.Fatalln(err)
	}
	defer productConn.Close()

	// Start server
	userGrpcServer.StartUserGrpcServer(host, grpcPort)
	userGrpcServer.StartGrpcGetwayServer(host, grpcGwPort)

}
