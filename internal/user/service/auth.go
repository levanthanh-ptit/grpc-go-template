package service

import (
	"grpc-go-templete/internal/user/domain"
	"grpc-go-templete/internal/user/provider"
	"grpc-go-templete/internal/user/repository"
)

type AuthService struct {
	passwordHashProvider provider.PasswordHashProvider
	tokenProvider        provider.TokenProvider

	userRepo repository.UserRepository
}

type LoginResult struct {
	Profile *domain.User
	Token   string
}

func NewAuthService(
	passwordHashProvider provider.PasswordHashProvider,
	tokenProvider provider.TokenProvider,
	userRepo repository.UserRepository,
) *AuthService {
	return &AuthService{
		passwordHashProvider: passwordHashProvider,
		tokenProvider:        tokenProvider,
		userRepo:             userRepo,
	}
}

func (s *AuthService) Login(username, password string) (*LoginResult, error) {
	user, err := s.userRepo.GetOne(&domain.User{Username: username})
	if err != nil {
		return nil, err
	}
	err = s.passwordHashProvider.ComparePassword(password, user.Password)
	if err != nil {
		return nil, err
	}
	token, err := s.tokenProvider.GenerateToken()
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		Profile: user,
		Token:   token,
	}, nil
}
