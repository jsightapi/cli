VERSION="0.0.1"
COMMIT_HASH=`git show -s --format=%H`
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
		-ldflags="-X 'main.Version=${VERSION}'" \
		-ldflags="-X 'main.CommitHash=${COMMIT_HASH}'" \
		-o ${BUILD_DIR}/jsight.so \
		.
