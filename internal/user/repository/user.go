package repository

import "grpc-go-templete/internal/user/domain"

type UserRepository interface {
	CreateIndexes()
	GetOne(query interface{}) (*domain.User, error)
	GetAll(query interface{}) ([]*domain.User, error)
	Count(query interface{}) (*int64, error)
	Create(entity *domain.User) (*domain.User, error)
	Update(query interface{}, update *domain.User) ([]*domain.User, error)
}
