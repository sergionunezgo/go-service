#!/bin/sh

echo "running docker container for golangci-lint"

docker run --rm -v $(pwd):/app -w /app \
    golangci/golangci-lint:v1.29.0 \
    /bin/bash -c "golangci-lint run -v --timeout 2m30s"