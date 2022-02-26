FROM homebrew/brew:3.3.16

WORKDIR /repo

RUN brew install go

RUN brew install protoc-gen-go

RUN brew install protoc-gen-go-grpc

RUN go get -u github.com/sergionunezgo/go-reuse/v2

ENTRYPOINT [ "script/protobuf/protoc.sh" ]