LINTER_VERSION=v1.54.0
GOPATH=$(shell go env GOPATH)

ENGINE_NAME ?= engine

GO_VERSION = $(shell go version | cut -d ' ' -f 3 | cut -c 3-10)
GIT_TAG = $(shell git describe --tags)
GIT_COMMIT = $(shell git describe --always --dirty)
NOW = $(shell TZ="GMT" LC_TIME="en_US.utf8" date)

BUILD_FLAGS = "\
-X 'main.GoVersion=$(GO_VERSION)'\
-X 'main.GitTag=$(GIT_TAG)'\
-X 'main.GitCommit=$(GIT_COMMIT)'\
-X 'main.Build=$(NOW)'"

## default: run clean, tidy, vendor, lint and build
default: clean tidy vendor lint build

## clean: clean the vendor and dist directories
clean:
	rm -rf vendor/
	rm -rf dist/

## tidy: add missing and remove unused modules
tidy:
	go mod tidy

## vendor: make vendored copy of dependencies
vendor:
	go mod vendor

## lint: download/install golangci-lint and analyse the source code with the configuration in .golangci.yml
lint:
	go fmt ./...
	command -v golangci-lint || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin $(LINTER_VERSION)
	golangci-lint run --timeout=5m

## build: build game engine to the dist directory
build:
	GOOS=js GOARCH=wasm go build -ldflags $(BUILD_FLAGS) -o dist/${ENGINE_NAME}.wasm ./cmd/wasm

## help: print this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
