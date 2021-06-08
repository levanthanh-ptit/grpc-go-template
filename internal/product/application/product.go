package application

import (
	"context"
	"grpc-go-templete/internal/product/helper"
	"grpc-go-templete/pkg/pb/product_pb"

	"github.com/levanthanh-ptit/go-ez/ez_grpc"
)

func (s *productsGrpcServer) CreateProduct(ctx context.Context, in *product_pb.Product) (*product_pb.Product, error) {
	err := in.Validate()
	if err != nil {
		return nil, ez_grpc.MakeInvalidArgument(err)
	}
	product, err := s.productService.CreateProduct(helper.ToProduct(in))
	if err != nil {
		return nil, err
	}
	return helper.ToProductGRPC(product), nil
}

func (s *productsGrpcServer) GetProduct(ctx context.Context, in *product_pb.ProductRequest) (*product_pb.ProductResponse, error) {
	return &product_pb.ProductResponse{Data: &product_pb.Product{}}, nil
}
