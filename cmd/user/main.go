package main

import (
	"grpc-go-templete/internal/user/api"
	"grpc-go-templete/internal/user/persistance"
	"grpc-go-templete/internal/user/service"
)

func main() {
	// Init Database
	mongoProvider, err := persistance.NewMongoProvider("")
	if err != nil {
		panic("DB not connect")
	}

	// Init repositories
	userRepo := persistance.NewUserPersistance(mongoProvider.GetDatabase("user"))

	// Init services
	userService := service.NewUserService(userRepo)

	// Init Application
	conn := api.InitUserGrpcServer(userService)
	api.InitGrpcGetway(conn)
}
