package service

import (
	"context"
	"grpc-go-templete/internal/product/domain"
	"grpc-go-templete/internal/product/repository"
)

type ProductService struct {
	userRepo repository.ProductRepository
}

func NewProductService(userRepo repository.ProductRepository) *ProductService {
	return &ProductService{
		userRepo,
	}
}

func (s ProductService) CreateProduct(ctx context.Context, user *domain.Product) (*domain.Product, error) {
	return s.userRepo.Create(ctx, user)
}

func (s ProductService) GetProductById(ctx context.Context, id domain.ID) (*domain.Product, error) {
	return s.userRepo.GetOne(ctx, domain.Product{ID: id})
}
