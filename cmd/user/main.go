package main

import (
	"fmt"
	"grpc-go-templete/internal/user/application"
	"grpc-go-templete/internal/user/infrastructure"
	"grpc-go-templete/internal/user/service"
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
	userRepo := infrastructure.NewUserRepositoryMongo(mongoProvider.GetDatabase("user_dev"))
	userRepo.CreateIndexes()

	// Init services
	userService := service.NewUserService(bcryptProvider, userRepo)
	authService := service.NewAuthService(bcryptProvider, jwtProvider, userRepo)

	// Application enums
	host := "localhost"
	grpcPort := "8081"
	grpcGwPort := "8091"

	// Init Application
	grpcServer := application.NewGrpcServer(host, grpcPort, userService, authService)
	grpcGetway := application.NewGrpcGetway(fmt.Sprintf("%s:%s", host, grpcPort), host, grpcGwPort)

	// Start GRPC server
	go grpcServer.ServerTCP()
	defer grpcGetway.Close()
	grpcGetway.Server()
}
