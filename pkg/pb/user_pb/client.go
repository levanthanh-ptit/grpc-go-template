package user_pb

import (
	"context"

	"google.golang.org/grpc"
)

func GetClient(adrress string, opts ...grpc.DialOption) (conn *grpc.ClientConn, client UsersClient, err error) {
	conn, err = grpc.DialContext(context.Background(), adrress, opts...)
	client = NewUsersClient(conn)
	return
}
