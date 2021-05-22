package application

import (
	"grpc-go-templete/pkg/pb/product_pb"

	"google.golang.org/grpc"
)

func (usersServer *usersGrpcServer) RegisterProductsClient(address string) (conn *grpc.ClientConn, err error) {
	conn, usersServer.ProductsClient, err = product_pb.GetClient(address, grpc.WithBlock(), grpc.WithInsecure())
	return
}
