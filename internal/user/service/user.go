package service

import (
	"grpc-go-templete/internal/user/model"
	"grpc-go-templete/internal/user/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo,
	}
}

func (s UserService) CreateUser(user *model.User) (*model.User, error) {
	return s.userRepo.Create(user)
}

func (s UserService) GetUserById(id model.ID) (*model.User, error) {
	return s.userRepo.GetOne(model.User{ID: id})
}
