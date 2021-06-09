package service

import (
	"context"
	"grpc-go-templete/internal/user/domain"
	"grpc-go-templete/internal/user/provider"
	"grpc-go-templete/internal/user/repository"
)

type AuthService struct {
	passwordHashProvider provider.PasswordHashProvider
	tokenProvider        provider.TokenProvider

	userRepo repository.UserRepository
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

type LoginResult struct {
	Profile *domain.User
	Token   string
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*LoginResult, error) {
	user, err := s.userRepo.GetOne(ctx, &domain.User{Username: username})
	if err != nil {
		return nil, err
	}
	err = s.passwordHashProvider.ComparePassword(password, user.Password)
	if err != nil {
		return nil, err
	}
	token, err := s.tokenProvider.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		Profile: user,
		Token:   token,
	}, nil
}

type VerifyTokenResult struct {
	Authenticated bool
	Payload       *domain.User
}

func (s *AuthService) VerifyToken(token string) (*VerifyTokenResult, error) {
	currentUser, err := s.tokenProvider.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	return &VerifyTokenResult{
		Authenticated: true,
		Payload:       currentUser.(*domain.User),
	}, nil
}
