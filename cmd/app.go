package cmd

import (
	"context"
	"fmt"
	"os"
	auth_services "template/internal/auth/services"
	user_repositories "template/internal/user/repositories"
	user_services "template/internal/user/services"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

type App struct {
	//DBs
	mongo *mongo.Database
	// Repositories
	userRepo user_repositories.UserRepo

	// Services
	userService *user_services.UserService

	basicAuthService *auth_services.BasicAuthService
}

func NewApp(ctx context.Context) (*App, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Minute)
	defer cancel()
	app := &App{}
	if err := app.connectDBs(ctx); err != nil {
		return nil, fmt.Errorf("app.connectDBs: %w", err)
	}
	if err := app.initRepositories(ctx); err != nil {
		return nil, fmt.Errorf("app.InitRepositories: %w", err)
	}
	app.initServices()
	return app, nil
}

func (a *App) connectDBs(ctx context.Context) (err error) {
	mongoCS, err := connstring.ParseAndValidate(os.Getenv("MONGO_DB"))
	if err != nil {
		return fmt.Errorf("connstring.ParseAndValidate: %w", err)
	}
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_DB")))
	if err != nil {
		return fmt.Errorf("mongo.Connect: %w", err)
	}
	a.mongo = mongoClient.Database(mongoCS.Database)
	return nil
}

func (a *App) initRepositories(ctx context.Context) error {
	a.userRepo = user_repositories.NewUserRepoMongo(a.mongo)
	return nil
}

func (a *App) initServices() {
	a.userService = user_services.NewUserService(a.userRepo)

	a.basicAuthService = auth_services.NewBasicAuthService(a.userService)
}

func (a *App) Start(ctx context.Context) {
	// TODO: start grpc
}
