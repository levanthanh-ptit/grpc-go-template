package helper

import (
	"grpc-go-templete/internal/product/model"
	"grpc-go-templete/pkg/pb/product_pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToUserGRPC(product *model.Product) *product_pb.Product {
	return &product_pb.Product{
		Id:   product.ID.Hex(),
		Name: *product.Name,
	}
}

func ToProduct(user *product_pb.Product) *model.Product {
	id, _ := primitive.ObjectIDFromHex(user.Id)
	return &model.Product{
		ID:   model.ID(id),
		Name: &user.Name,
	}
}
