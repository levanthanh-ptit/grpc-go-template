package service

import (
	"grpc-go-templete/internal/user/domain"
	"grpc-go-templete/internal/user/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s UserService) CreateUser(user *domain.User) (*domain.User, error) {
	return s.userRepo.Create(user)
}

func (s UserService) GetUserById(id domain.ID) (*domain.User, error) {
	return s.userRepo.GetOne(domain.User{ID: id})
}
