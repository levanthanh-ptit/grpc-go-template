package application

import (
	"context"
	"grpc-go-templete/pkg/pb/product_pb"
)

func (s *productsGrpcServer) GetProduct(ctx context.Context, in *product_pb.ProductRequest) (*product_pb.ProductResponse, error) {
	return &product_pb.ProductResponse{Data: &product_pb.Product{}}, nil
}
