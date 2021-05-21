package repository

import "grpc-go-templete/internal/product/model"

type ProductRepository interface {
	GetOne(query interface{}) (*model.Product, error)
	GetAll(query, options interface{}) ([]*model.Product, error)
	Count(query interface{}) (*int64, error)
	Create(entity *model.Product) (*model.Product, error)
	Update(query interface{}, update *model.Product, options interface{}) ([]*model.Product, error)
}
