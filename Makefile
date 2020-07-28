
run: build
	./build/gorest

build:
	./script/build.sh
.PHONY: build

format:
	./script/format.sh

lint:
	./script/lint.sh

test:
	./script/test.sh
