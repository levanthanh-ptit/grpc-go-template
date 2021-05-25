package service

import (
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

func (s ProductService) CreateProduct(user *domain.Product) (*domain.Product, error) {
	return s.userRepo.Create(user)
}

func (s ProductService) GetProductById(id domain.ID) (*domain.Product, error) {
	return s.userRepo.GetOne(domain.Product{ID: id})
}
