package application

import (
	"context"
	"grpc-go-templete/pkg/pb/user_pb"
)

func (s *usersGrpcServer) Login(ctx context.Context, in *user_pb.LoginRequest) (*user_pb.LoginResponse, error) {
	return &user_pb.LoginResponse{}, nil
}
