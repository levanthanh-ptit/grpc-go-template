package main

import (
	"grpc-go-templete/internal/auth/application"
	"grpc-go-templete/internal/auth/persistance"
	"grpc-go-templete/internal/auth/service"
	"grpc-go-templete/pkg/provider"
)

func main() {
	// Init Database
	mongoProvider, err := provider.NewMongoProvider("mongodb://localhost:27017")
	if err != nil {
		panic("DB not connect")
	}

	// Init repositories
	accountRepo := persistance.NewAccountPersistance(mongoProvider.GetDatabase("auth_dev"))

	// Init services
	authService := service.NewAuthService(accountRepo)

	// Application enums
	host := "localhost"
	grpcPort := "8082"

	// Init Application
	authGrpcServer := application.NewAuthGrpcServer(host, grpcPort, authService)

	// Start GRPC server
	authGrpcServer.StartGrpcServer()
}
