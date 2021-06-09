package helper

import (
	"grpc-go-templete/internal/product/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FromRawID(in interface{}) (domain.ID, error) {
	return primitive.ObjectIDFromHex(in.(string))
}
