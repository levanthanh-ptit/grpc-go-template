package application

import (
	"context"
	"grpc-go-templete/internal/user/helper"
	"grpc-go-templete/pkg/grpc_utils"
	"grpc-go-templete/pkg/pb/user_pb"
)

func (s *usersGrpcServer) CreateUser(ctx context.Context, in *user_pb.User) (*user_pb.User, error) {
	err := in.Validate()
	if err != nil {
		return nil, grpc_utils.MakeInvalidArgument(err)
	}
	user, err := s.userService.CreateUser(helper.ToUser(in))
	if err != nil {
		return nil, err
	}
	return helper.ToUserGRPC(user), nil
}

func (s *usersGrpcServer) GetUser(ctx context.Context, in *user_pb.UserRequest) (*user_pb.UserResponse, error) {
	resp := &user_pb.UserResponse{}
	return resp, nil
}

func (s *usersGrpcServer) GetUserList(ctx context.Context, in *user_pb.UserListRequest) (*user_pb.UserListResponse, error) {
	resp := &user_pb.UserListResponse{}
	return resp, nil
}
