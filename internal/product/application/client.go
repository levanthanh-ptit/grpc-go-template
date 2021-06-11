package application

import (
	"grpc-go-templete/pkg/pb/user_pb"
	"log"

	"google.golang.org/grpc"
)

// RegisterUsersClient register user client
func (s *GrpcGetwayServer) RegisterUsersClient(address string) (conn *grpc.ClientConn, err error) {
	conn, s.usersClient, err = user_pb.GetClient(address, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("User service client registered.")
	return
}
