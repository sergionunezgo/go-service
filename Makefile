
run: build
	./build/goservice

build:
	./script/build.sh
.PHONY: build

format:
	./script/format.sh

lint:
	./script/lint.sh

test:
	./script/test.sh
