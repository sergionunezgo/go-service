#!/bin/sh

PROTO_DIR=goservice/greeting

# We need to point the proto compiler to a direction in the file system where to find: import "goreuse/test/test.proto";
# We can use go list to extract the path to a downloaded module.
GO_REUSE_PATH=$(go list -m -f "{{.Dir}}" github.com/sergionunezgo/go-reuse)
GO_REUSE_V2_PATH=$(go list -m -f "{{.Dir}}" github.com/sergionunezgo/go-reuse/v2)

echo "GO_REUSE_PATH=${GO_REUSE_PATH}"
echo "GO_REUSE_V2_PATH=${GO_REUSE_V2_PATH}"

protoc --proto_path=${PROTO_DIR} --go_out=. --go-grpc_out=. \
    -I"${GO_REUSE_PATH}" \
    -I"${GO_REUSE_V2_PATH}" \
    --go_opt=module=github.com/sergionunezgo/go-service \
    --go-grpc_opt=module=github.com/sergionunezgo/go-service \
    greeting.proto