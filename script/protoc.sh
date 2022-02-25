#!/bin/sh

PROTO_DIR=goservice/greeting

# We need to point the proto compiler to a direction in the file system where to find: import "goreuse/test/test.proto";
GO_REUSE_PATH="/Users/sergio/GIT/github.com/sergionunezgo/go-reuse"

protoc --proto_path=${PROTO_DIR} --go_out=. --go-grpc_out=. \
    -I"${GO_REUSE_PATH}" \
    --go_opt=module=github.com/sergionunezgo/go-service \
    --go-grpc_opt=module=github.com/sergionunezgo/go-service \
    greeting.proto