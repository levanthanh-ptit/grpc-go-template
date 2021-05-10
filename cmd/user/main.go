package main

import (
	"grpc-go-templete/internal/user/api"
)

func main() {
	conn := api.InitUserGrpcServer()
	api.InitGrpcGetway(conn)
}
