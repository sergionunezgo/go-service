#!/bin/sh

ENV_FILE=.env.list

source ${ENV_FILE}

echo "using env vars from ${ENV_FILE}"
echo "running docker container ${DOCKER_REPO}"

docker run -it --rm --read-only -p ${API_PORT}:${API_PORT} \
    --env-file $(pwd)/${ENV_FILE} ${DOCKER_REPO} $@