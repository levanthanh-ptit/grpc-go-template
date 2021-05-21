package repository

import "grpc-go-templete/internal/user/model"

type UserRepository interface {
	GetOne(query interface{}) (*model.User, error)
	GetAll(query, options interface{}) ([]*model.User, error)
	Count(query interface{}) (*int64, error)
	Create(entity *model.User) (*model.User, error)
	Update(query interface{}, update *model.User, options interface{}) ([]*model.User, error)
}
