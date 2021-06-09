package service

import (
	"context"
	"grpc-go-templete/internal/user/domain"
	"grpc-go-templete/internal/user/provider"
	"grpc-go-templete/internal/user/repository"
)

// UserService define service for User
type UserService struct {
	passwordHashProvider provider.PasswordHashProvider
	userRepo             repository.UserRepository
}

// NewUserService create new UserService instance
func NewUserService(
	passwordHashProvider provider.PasswordHashProvider,
	userRepo repository.UserRepository,
) *UserService {
	return &UserService{
		passwordHashProvider: passwordHashProvider,
		userRepo:             userRepo,
	}
}

// CreateUser method for create new User
func (s UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// Hash the password
	pwd, err := s.passwordHashProvider.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = pwd
	return s.userRepo.Create(ctx, user)
}

// GetUserByID method for get one User
func (s UserService) GetUserByID(ctx context.Context, id domain.ID) (*domain.User, error) {
	return s.userRepo.GetOne(ctx, domain.User{ID: id})
}

// GetUserList method for get Users
func (s UserService) GetUserList(ctx context.Context, query *domain.User, limit, offset int64) ([]*domain.User, error) {
	return s.userRepo.GetAll(ctx, query, limit, offset)
}

// UpdateUsers method for update User(s)
func (s UserService) UpdateUsers(ctx context.Context, query, update *domain.User, limit, offset int64) ([]*domain.User, error) {
	return s.userRepo.Update(ctx, query, update, limit, offset)
}
