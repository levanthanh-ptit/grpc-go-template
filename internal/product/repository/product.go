package repository

import (
	"context"
	"grpc-go-templete/internal/product/domain"
)

type ProductRepository interface {
	GetOne(ctx context.Context, query interface{}) (*domain.Product, error)
	GetAll(ctx context.Context, query interface{}, limit, offset int64) ([]*domain.Product, error)
	Count(ctx context.Context, query interface{}) (*int64, error)
	Create(ctx context.Context, entity *domain.Product) (*domain.Product, error)
	Update(ctx context.Context, query interface{}, update *domain.Product, limit, offset int64) ([]*domain.Product, error)
}
