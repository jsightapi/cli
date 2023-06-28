VERSION="0.0.1"
COMMIT_HASH=`git show -s --format=%H`

.PHONY: all
all: fmt lint build

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
  go build \
		-ldflags="-X 'main.Version=${VERSION}'" \
		-ldflags="-X 'main.CommitHash=${COMMIT_HASH}'" \
		-o build/jsight-cli \
		.
