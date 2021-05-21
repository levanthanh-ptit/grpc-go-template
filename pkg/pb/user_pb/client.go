package user_pb

import "google.golang.org/grpc"

func GetClient(adrress string, opts ...grpc.DialOption) (conn *grpc.ClientConn, client UsersClient, err error) {
	conn, err = grpc.Dial(adrress, opts...)
	client = NewUsersClient(conn)
	return
}
