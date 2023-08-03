FROM golang:1.18-stretch AS build-stage
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
      | sh -s -- -b $(go env GOPATH)/bin v1.53.3

COPY . /go/src/github.com/jsightapi/cli
WORKDIR /go/src/github.com/jsightapi/cli/docker
RUN make

FROM scratch AS export-stage
COPY --from=build-stage /tmp/build .
