package main

import (
	"grpc-go-templete/internal/user/application"
	"grpc-go-templete/internal/user/infrastructure"
	"grpc-go-templete/internal/user/service"
	"log"
	"time"

	"github.com/levanthanh-ptit/go-ez/ez_provider"
)

func main() {
	// Init Database
	mongoProvider, err := ez_provider.NewMongoProvider("mongodb://localhost:27017")
	if err != nil {
		panic("DB not connect")
	}

	// Init providers
	bcryptProvider := ez_provider.NewBcryptProvider("salt", 10)
	jwtSigningKey := "strongJwTKeY"
	jwtAccessTokenDuration := 5 * time.Minute
	jwtRefreshTokenDuration := 24 * time.Hour
	jwtProvider := ez_provider.NewJWTProvider(
		jwtSigningKey,
		jwtAccessTokenDuration,
		jwtRefreshTokenDuration,
		infrastructure.AddClaims,
		infrastructure.ExtractClaims,
	)

	// Init repositories
	userRepo := infrastructure.NewUserPersistance(mongoProvider.GetDatabase("user_dev"))
	userRepo.CreateIndexes()

	// Init services
	userService := service.NewUserService(bcryptProvider, userRepo)
	authService := service.NewAuthService(bcryptProvider, jwtProvider, userRepo)

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
