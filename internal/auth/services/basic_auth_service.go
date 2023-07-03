package services

import (
	"context"
	"fmt"
	auth_dtos "template/internal/auth/dtos"
	user_dtos "template/internal/user/dtos"
	user_entities "template/internal/user/entities"
	user_services "template/internal/user/services"
)

type BasicAuthService struct {
	userService *user_services.UserService
}

func NewBasicAuthService(userService *user_services.UserService) *BasicAuthService {
	return &BasicAuthService{
		userService: userService,
	}
}

func (s *BasicAuthService) SignUp(ctx context.Context, data *auth_dtos.SignUpDto) (*user_entities.User, error) {
	user, err := s.userService.Create(ctx, &user_dtos.CreateUserDto{
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("s.userService.Create: %w", err)
	}
	return user, nil
}
