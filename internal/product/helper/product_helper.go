package helper

import (
	"grpc-go-templete/internal/product/domain"
	"grpc-go-templete/pkg/pb/product_pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToProductGRPC(product *domain.Product) *product_pb.Product {
	return &product_pb.Product{
		Id:   product.ID.Hex(),
		Name: *product.Name,
	}
}

func ToProduct(user *product_pb.Product) *domain.Product {
	id, _ := primitive.ObjectIDFromHex(user.Id)
	return &domain.Product{
		ID:   domain.ID(id),
		Name: &user.Name,
	}
}
