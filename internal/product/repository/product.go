package repository

import "grpc-go-templete/internal/product/domain"

type ProductRepository interface {
	GetOne(query interface{}) (*domain.Product, error)
	GetAll(query, options interface{}) ([]*domain.Product, error)
	Count(query interface{}) (*int64, error)
	Create(entity *domain.Product) (*domain.Product, error)
	Update(query interface{}, update *domain.Product, options interface{}) ([]*domain.Product, error)
}
