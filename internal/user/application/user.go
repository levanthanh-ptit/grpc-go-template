package application

import (
	"context"
	"grpc-go-templete/internal/user/helper"
	"grpc-go-templete/pkg/pb/user_pb"
)

func (s *userGrpcServer) CreateUser(ctx context.Context, in *user_pb.User) (*user_pb.User, error) {
	user, err := s.userService.CreateUser(helper.ToUser(in))
	if err != nil {
		return nil, err
	}
	resp := helper.ToUserGRPC(user)
	return resp, nil
}

func (s *userGrpcServer) GetUser(ctx context.Context, in *user_pb.UserRequest) (*user_pb.UserResponse, error) {
	resp := &user_pb.UserResponse{}
	return resp, nil
}

func (s *userGrpcServer) GetUserList(ctx context.Context, in *user_pb.UserListRequest) (*user_pb.UserListResponse, error) {
	resp := &user_pb.UserListResponse{}
	return resp, nil
}
