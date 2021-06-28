run: build
	./build/go-service

.PHONY: build
build:
	./script/build.sh

run-docker: build-docker
	./script/run-docker.sh

build-docker:
	./script/build-docker.sh

format:
	./script/format.sh

lint:
	./script/lint.sh

test:
	./script/test.sh

precommit: format lint build test