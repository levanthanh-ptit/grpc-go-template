package repository

import (
	"context"
	"grpc-go-templete/internal/user/domain"
)

type UserRepository interface {
	CreateIndexes()
	GetOne(ctx context.Context, query interface{}) (*domain.User, error)
	GetAll(ctx context.Context, query interface{}, limit, offset int64) ([]*domain.User, error)
	Count(ctx context.Context, query interface{}) (*int64, error)
	Create(ctx context.Context, entity *domain.User) (*domain.User, error)
	Update(ctx context.Context, query interface{}, update *domain.User, limit, offset int64) ([]*domain.User, error)
}
