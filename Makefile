VERSION="0.0.1"
COMMIT_HASH=`git show -s --format=%H`
WORK_DIR="/go/src/github.com/jsightapi/cli"

.PHONY: all
all: build test

.PHONY: build
build: 
	docker build --rm -f "docker/Dockerfile" --progress plain --output build .

.PHONY: test
test:
	docker compose -f "docker/docker-compose.yml" up

.PHONY: fmt
fmt:
	go fmt ./...

.PHONE: lint
lint:
	golangci-lint run

.PHONY: dev
dev: fmt lint test