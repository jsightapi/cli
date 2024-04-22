WORK_DIR=/go/src/github.com/jsightapi/cli
BUILD_DIR=/tmp/build

.PHONY: all
all: fmt lint build

.PHONY: fmt
fmt:
	cd ${WORK_DIR} && go fmt ./...

.PHONY: lint
lint:
	cd ${WORK_DIR} && golangci-lint run

.PHONY: build
build:
	cd ${WORK_DIR} && go build \
		-o ${BUILD_DIR}/jsight \
		.
