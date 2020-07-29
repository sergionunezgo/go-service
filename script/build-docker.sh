#!/bin/sh

echo "running docker image build ${DOCKER_REPO}"

docker build -t ${DOCKER_REPO}:latest --build-arg VERSION=${VERSION} .