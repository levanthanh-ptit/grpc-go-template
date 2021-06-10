package application

import (
	"context"
	"grpc-go-templete/internal/user/helper"
	"grpc-go-templete/pkg/pb/user_pb"
)

func (s *GrpcServer) Login(ctx context.Context, in *user_pb.LoginRequest) (*user_pb.LoginResponse, error) {
	result, err := s.authService.Login(ctx, in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	return &user_pb.LoginResponse{
		Profile: helper.ToUserGRPC(result.Profile),
		Token:   result.Token,
	}, nil
}

func (s *GrpcServer) VerifyAuthToken(ctx context.Context, in *user_pb.VerifyAuthTokenRequest) (*user_pb.VerifyAuthTokenResponse, error) {
	result, err := s.authService.VerifyToken(in.Token)
	if err != nil {
		return nil, err
	}
	return &user_pb.VerifyAuthTokenResponse{
		Authenticated: result.Authenticated,
		Payload:       helper.ToUserGRPC(result.Payload),
	}, nil
}
