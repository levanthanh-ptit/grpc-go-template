package application

import (
	"grpc-go-templete/pkg/pb/product_pb"
	"log"

	"google.golang.org/grpc"
)

func (s *GrpcServer) RegisterProductsClient(address string) (err error) {
	s.productConn, s.productClient, err = product_pb.GetClient(address, grpc.WithInsecure())
	if err != nil {
		return
	}
	log.Println("Product service client registered.")
	return
}
