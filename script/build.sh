#!/bin/sh

echo "running go build"

CGO_ENABLED=0 go build -o build/go-service ./cmd/go-service