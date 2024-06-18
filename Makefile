GO := go
GOCOVER := $(GO) tool cover
GOTEST := $(GO) test
BINARY_NAME := objectid
VERSION := $(shell git describe --tags)
BUILD := $(shell date +%FT%T%z)


.PHONY: all
all: test build

.PHONY: build
build: bin/objectid

bin/objectid:
	@cd cmd/objectid && GOARCH=amd64 GOOS=linux go build ${BUILD_ARGS} -ldflags "-w -s -X main.version=${VERSION} -X main.build=${BUILD}" -o ../../$@

.PHONY: clean
clean:
	@rm -rf bin/*

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: test/cover
test/cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out
