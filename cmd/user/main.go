package main

import (
	"grpc-go-templete/internal/user/application"
	"grpc-go-templete/internal/user/persistance"
	"grpc-go-templete/internal/user/service"
	"grpc-go-templete/pkg/provider"
	"log"
)

func main() {
	// Init Database
	mongoProvider, err := provider.NewMongoProvider("mongodb://localhost:27017")
	if err != nil {
		panic("DB not connect")
	}

	// Init providers
	bcryptProvider := provider.NewBcryptProvider("salt", 10)

	// Init repositories
	userRepo := persistance.NewUserPersistance(mongoProvider.GetDatabase("user_dev"))

	// Init services
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(bcryptProvider, userRepo)

	// Application enums
	host := "localhost"
	grpcPort := "8081"
	grpcGwPort := "8091"
	productServerAddress := "localhost:8082"

	// Init Application
	userGrpcServer := application.NewUserGrpcServer(host, grpcPort, userService, authService)

	// Start GRPC server
	go userGrpcServer.StartUserGrpcServer()

	// Init Clients
	productConn, err := userGrpcServer.RegisterProductsClient(productServerAddress)
	if err != nil {
		log.Fatalln(err)
	}

	var closeClients = func() {
		productConn.Close()
	}
	defer closeClients()

	// Start GRPC Getway
	userGrpcServer.StartGrpcGetwayServer(host, grpcGwPort)

}
