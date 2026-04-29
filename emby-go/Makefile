.PHONY: build run test clean deps

MODULE_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BINARY_NAME := emby-server
BUILD_DIR := $(MODULE_DIR)/bin

build:
	@cd $(MODULE_DIR) && go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/emby-server/

run: build
	@cd $(MODULE_DIR) && $(BUILD_DIR)/$(BINARY_NAME)

test:
	@cd $(MODULE_DIR) && go test -v ./...

clean:
	@cd $(MODULE_DIR) && rm -rf $(BUILD_DIR)
	@cd $(MODULE_DIR) && go clean

deps:
	@cd $(MODULE_DIR) && go mod tidy

lint:
	@cd $(MODULE_DIR) && golangci-lint run

format:
	@cd $(MODULE_DIR) && gofmt -s -w .
