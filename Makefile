GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test


.PHONY: all
all: test build

.PHONY: build
build: bin/objectid

bin/objectid:
	$(GO) build -o bin/objectid cmd/objectid/main.go

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: test/cover
test/cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out
