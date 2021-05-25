package helper

import (
	"grpc-go-templete/internal/user/domain"
	"grpc-go-templete/pkg/pb/user_pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToUserGRPC(user *domain.User) *user_pb.User {
	return &user_pb.User{
		Id:       user.ID.Hex(),
		Name:     *user.Name,
		Username: user.Username,
		Old:      *user.Old,
	}
}

func ToUser(user *user_pb.User) *domain.User {
	id, _ := primitive.ObjectIDFromHex(user.Id)
	return &domain.User{
		ID:       domain.ID(id),
		Name:     &user.Name,
		Username: user.Username,
		Password: user.Password,
		Old:      &user.Old,
	}
}
