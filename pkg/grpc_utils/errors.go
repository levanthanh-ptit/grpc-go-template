package grpc_utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MakeInvalidArgument(err error) error {
	return status.Error(codes.InvalidArgument, err.Error())
}

func MakeAlreadyExists(err error) error {
	return status.Error(codes.AlreadyExists, err.Error())
}

func MakeDeadlineExceeded(err error) error {
	return status.Error(codes.DeadlineExceeded, err.Error())
}

func MakePermissionDenied(err error) error {
	return status.Error(codes.PermissionDenied, err.Error())
}

func MakeUnauthenticated(err error) error {
	return status.Error(codes.Unauthenticated, err.Error())
}

func MakeNotFound(err error) error {
	return status.Error(codes.NotFound, err.Error())
}
