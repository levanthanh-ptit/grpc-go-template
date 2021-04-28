#!/bin/bash
protoc -I ./api/proto \
   --go_out ./pkg/pb \
   --go_opt paths=source_relative \
   --go-grpc_out ./pkg/pb \
   --go-grpc_opt paths=source_relative \
   --grpc-gateway_out ./pkg/pb \
   --grpc-gateway_opt paths=source_relative \
   --grpc-gateway_opt logtostderr=true \
   --grpc-gateway_opt generate_unbound_methods=true \
   $(find ./api/proto -iwholename "*_pb/*.proto")