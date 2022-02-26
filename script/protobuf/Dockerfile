FROM homebrew/brew:3.3.16

RUN brew install go

RUN brew install protoc-gen-go

RUN brew install protoc-gen-go-grpc

RUN go get -u github.com/sergionunezgo/go-reuse/v2

RUN go get -u github.com/sergionunezgo/go-reuse

WORKDIR /repo

ENTRYPOINT [ "script/protobuf/protoc.sh" ]