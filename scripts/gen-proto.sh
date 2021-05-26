#!/bin/bash
## Find all proto files in projects.
proto_files=$(find ./api/proto -iwholename "*_pb/*.proto")
## Get protoc-genvalidate folder of installed pkg.
protoc_gen_validate_path=$(find ${GOPATH}/pkg/mod/github.com/envoyproxy -type d -iname "protoc-gen-validate*")
protoc -I ./api/proto \
   --go_out ./pkg/pb \
   --go_opt paths=source_relative \
   --go-grpc_out ./pkg/pb \
   --go-grpc_opt paths=source_relative \
   --grpc-gateway_out ./pkg/pb \
   --grpc-gateway_opt paths=source_relative \
   --grpc-gateway_opt logtostderr=true \
   --grpc-gateway_opt generate_unbound_methods=true \
   -I $protoc_gen_validate_path \
   --validate_out ./pkg/pb \
   --validate_opt lang=go \
   --validate_opt paths=source_relative \
   $proto_files