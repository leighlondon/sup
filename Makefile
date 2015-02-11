# Makefile for common operations.
.PHONY: build fmt lint test
all: build fmt lint test
build:
	go build
fmt:
	go fmt
clean:
	go clean
lint:
	golint
test:
	go test
