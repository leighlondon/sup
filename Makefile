# Makefile for common operations.
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
