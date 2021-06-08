package application

import (
	"grpc-go-templete/pkg/pb/product_pb"
	"log"

	"google.golang.org/grpc"
)

func (usersServer *usersGrpcServer) RegisterProductsClient(address string) (conn *grpc.ClientConn, err error) {
	conn, usersServer.ProductsClient, err = product_pb.GetClient(address, grpc.WithInsecure())
	log.Println("Product service client registered.")
	return
}
