package main

import (
	"grpc-go-templete/internal/user/api"
	"grpc-go-templete/internal/user/persistance"
	"grpc-go-templete/internal/user/service"
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

	host := "localhost"
	grpcPort := "8081"
	grpcGwPort := "8091"

	// Init Application
	conn := api.InitUserGrpcServer(host, grpcPort, userService)
	api.InitGrpcGetway(host, grpcGwPort, conn)
}
