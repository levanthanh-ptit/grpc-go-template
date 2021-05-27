package service

import (
	"grpc-go-templete/internal/user/domain"
	"grpc-go-templete/internal/user/provider"
	"grpc-go-templete/internal/user/repository"
)

type UserService struct {
	passwordHashProvider provider.PasswordHashProvider
	userRepo             repository.UserRepository
}

func NewUserService(
	passwordHashProvider provider.PasswordHashProvider,
	userRepo repository.UserRepository,
) *UserService {
	return &UserService{
		passwordHashProvider: passwordHashProvider,
		userRepo:             userRepo,
	}
}

func (s UserService) CreateUser(user *domain.User) (*domain.User, error) {
	// Hash the password
	pwd, err := s.passwordHashProvider.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = pwd
	return s.userRepo.Create(user)
}

func (s UserService) GetUserById(id domain.ID) (*domain.User, error) {
	return s.userRepo.GetOne(domain.User{ID: id})
}
