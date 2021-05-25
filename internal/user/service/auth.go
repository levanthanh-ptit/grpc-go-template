package service

import (
	"grpc-go-templete/internal/user/domain"
	"grpc-go-templete/internal/user/repository"
	"grpc-go-templete/pkg/provider"
)

type AuthService struct {
	bcryptProvider *provider.BcryptProvider

	userRepo repository.UserRepository
}

func NewAuthService(bcryptProvider *provider.BcryptProvider, userRepo repository.UserRepository) *AuthService {
	return &AuthService{
		bcryptProvider: bcryptProvider,
		userRepo:       userRepo,
	}
}

func (s *AuthService) Login(username, password string) error {
	user, err := s.userRepo.GetOne(&domain.User{Username: username})
	if err != nil {
		return err
	}
	err = s.bcryptProvider.ComparePassword(password, user.Password)
	if err != nil {
		return err
	}
	return nil
}
