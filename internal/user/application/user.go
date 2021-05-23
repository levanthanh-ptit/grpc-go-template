package application

import (
	"context"
	"grpc-go-templete/internal/user/helper"
	"grpc-go-templete/pkg/pb/product_pb"
	"grpc-go-templete/pkg/pb/user_pb"
)

func (s *usersGrpcServer) CreateUser(ctx context.Context, in *user_pb.User) (*user_pb.User, error) {
	user, err := s.userService.CreateUser(helper.ToUser(in))
	if err != nil {
		return nil, err
	}
	resp := helper.ToUserGRPC(user)
	_, err = s.ProductsClient.CreateProduct(ctx, &product_pb.Product{
		Name: *user.Name,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *usersGrpcServer) GetUser(ctx context.Context, in *user_pb.UserRequest) (*user_pb.UserResponse, error) {
	resp := &user_pb.UserResponse{}
	return resp, nil
}

func (s *usersGrpcServer) GetUserList(ctx context.Context, in *user_pb.UserListRequest) (*user_pb.UserListResponse, error) {
	resp := &user_pb.UserListResponse{}
	return resp, nil
}
