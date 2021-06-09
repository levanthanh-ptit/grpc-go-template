package application

import (
	"context"
	"grpc-go-templete/internal/user/helper"
	"grpc-go-templete/pkg/pb/user_pb"

	"github.com/levanthanh-ptit/go-ez/ez_grpc"
)

func (s *usersGrpcServer) CreateUser(ctx context.Context, in *user_pb.User) (*user_pb.User, error) {
	err := in.Validate()
	if err != nil {
		return nil, ez_grpc.MakeInvalidArgument(err)
	}
	user, err := s.userService.CreateUser(ctx, helper.ToUser(in))
	if err != nil {
		return nil, err
	}
	return helper.ToUserGRPC(user), nil
}

func (s *usersGrpcServer) GetUser(ctx context.Context, in *user_pb.UserRequest) (*user_pb.UserResponse, error) {
	userID, err := helper.FromRawID(in.Id)
	if err != nil {
		return nil, ez_grpc.MakeInvalidArgument(err)
	}
	userList, err := s.userService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, ez_grpc.MakeInvalidArgument(err)
	}
	return &user_pb.UserResponse{
		Data: helper.ToUserGRPC(userList),
	}, nil
}

func (s *usersGrpcServer) GetUserList(ctx context.Context, in *user_pb.UserListRequest) (*user_pb.UserListResponse, error) {
	resp := &user_pb.UserListResponse{}
	return resp, nil
}
