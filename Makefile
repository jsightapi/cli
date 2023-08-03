VERSION="0.0.1"
COMMIT_HASH=`git show -s --format=%H`
WORK_DIR="/go/src/github.com/jsightapi/cli"

.PHONY: all
all: build

.PHONY: build
build: 
	docker build --rm -f "docker/Dockerfile" --output build .
