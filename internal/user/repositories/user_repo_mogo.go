package repositories

import (
	"context"
	"fmt"
	"template/internal/user/dtos"
	"template/internal/user/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoMongo struct {
	db        *mongo.Database
	usersColl *mongo.Collection
}

func NewUserRepoMongo(db *mongo.Database) UserRepo {
	return &UserRepoMongo{
		db:        db,
		usersColl: db.Collection("users"),
	}
}

func (r *UserRepoMongo) Create(ctx context.Context, data *dtos.CreateUserDto) (*entities.User, error) {
	result, err := r.usersColl.InsertOne(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("r.usersColl.InsertOne: %w", err)
	}
	return &entities.User{
		ID:       result.InsertedID.(primitive.ObjectID),
		Email:    data.Email,
		Password: data.Password,
	}, nil
}
