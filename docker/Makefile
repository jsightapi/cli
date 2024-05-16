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
		-o ${BUILD_DIR}/linux-x64/jsight \
		.

	cd ${WORK_DIR} && GOOS=darwin GOARCH=amd64 go build \
		-o ${BUILD_DIR}/macos-x64/jsight \
		.

	cd ${WORK_DIR} && GOOS=windows GOARCH=amd64 go build \
		-o ${BUILD_DIR}/windows-x64/jsight.exe \
		.
	
