package product_pb

import (
	"context"

	"google.golang.org/grpc"
)

func GetClient(adrress string, opts ...grpc.DialOption) (conn *grpc.ClientConn, client ProductsClient, err error) {
	conn, err = grpc.DialContext(context.Background(), adrress, opts...)
	client = NewProductsClient(conn)
	return
}
