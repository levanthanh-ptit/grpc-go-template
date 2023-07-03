package services

import (
	"context"
	"fmt"
	"template/internal/user/dtos"
	"template/internal/user/entities"
	"template/internal/user/repositories"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Create(ctx context.Context, data *dtos.CreateUserDto) (*entities.User, error) {
	user, err := s.userRepo.Create(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("s.userRepo.Create: %w", err)
	}
	return user, nil
}
