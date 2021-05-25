package service

import (
	"grpc-go-templete/internal/auth/repository"
)

type AuthService struct {
	accountRepo repository.AccountRepository
}

func NewAuthService(accountRepo repository.AccountRepository) *AuthService {
	return &AuthService{
		accountRepo,
	}
}
