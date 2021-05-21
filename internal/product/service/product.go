package service

import (
	"grpc-go-templete/internal/product/model"
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

func (s ProductService) CreateProduct(user *model.Product) (*model.Product, error) {
	return s.userRepo.Create(user)
}

func (s ProductService) GetProductById(id model.ID) (*model.Product, error) {
	return s.userRepo.GetOne(model.Product{ID: id})
}
