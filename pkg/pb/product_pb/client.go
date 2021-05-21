package product_pb

import "google.golang.org/grpc"

func GetClient(adrress string, opts ...grpc.DialOption) (conn *grpc.ClientConn, client ProductsClient, err error) {
	conn, err = grpc.Dial(adrress, opts...)
	client = NewProductsClient(conn)
	return
}
