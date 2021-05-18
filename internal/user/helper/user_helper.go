package helper

import (
	"grpc-go-templete/internal/user/model"
	"grpc-go-templete/pkg/pb/user_pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToUserGRPC(user *model.User) *user_pb.User {
	return &user_pb.User{
		Id:       user.ID.ToString(),
		Name:     "",
		Username: "",
		Old:      0,
	}
}

func ToUser(user *user_pb.User) *model.User {
	id, _ := primitive.ObjectIDFromHex(user.Id)
	return &model.User{
		ID:   model.ID(id),
		Name: &user.Name,
	}
}
