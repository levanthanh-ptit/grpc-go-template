package repositories

import (
	"context"
	"template/internal/user/dtos"
	"template/internal/user/entities"
)

type UserRepo interface {
	Create(context.Context, *dtos.CreateUserDto) (*entities.User, error)
}
