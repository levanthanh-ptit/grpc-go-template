package grpc

import (
	"context"
	basic_authv1 "template/pkg/proto/auth/basic_auth/v1"
)

type BasicAuthService struct {
	basic_authv1.UnimplementedBasicAuthServiceServer
}

func (*BasicAuthService) SignUp(ctx context.Context, data *basic_authv1.SignUpRequest) (*basic_authv1.SignUpResponse, error) {
	return &basic_authv1.SignUpResponse{}, nil
}
