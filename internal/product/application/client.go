package application

import (
	"grpc-go-templete/pkg/pb/user_pb"
	"log"

	"google.golang.org/grpc"
)

func (productsServer *productsGrpcServer) RegisterUsersClient(address string) (conn *grpc.ClientConn, err error) {
	conn, productsServer.UsersClient, err = user_pb.GetClient(address, grpc.WithInsecure())
	log.Println("User service client registered.")
	return
}
