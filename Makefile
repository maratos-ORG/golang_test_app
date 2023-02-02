.DEFAULT_GOAL := help

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

ifeq ($(GOPATH),)
GOPATH := ~/go
endif

APPNAME = golang_test_app
TAG=$(shell git describe --tags |cut -d- -f1)
COMMIT=$(shell git rev-parse --short HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
PROJECT_DIR = cmd
RELEASE ?= dev
BINARY_NAME ?= bin/golang_test_app
BUILD_TIME ?= $(shell date '+%Y-%m-%d_%H:%M:%S')
COMMIT_HASH = $(shell git rev-parse --short HEAD)

.PHONY: all
all: linter test build ## Run linter, tests and build a package

dep: ## Get the dependencies
	go mod download

test: ## Run tests
	go test -race -coverprofile=coverage.out ./$(PROJECT_DIR)/...

linter: ## Apply linter
	golangci-lint run --timeout 5m -E golint -e '(method|func) [a-zA-Z]+ should be [a-zA-Z]+'

clean: ## Clean before build
	@go clean ./...

build: clean ## Build package
	@go build -ldflags "-s -w -X main.appName=${APPNAME} -X main.gitTag=${TAG} -X main.gitCommit=${COMMIT} -X main.gitBranch=${BRANCH}'" -o ${BINARY_NAME} ./$(PROJECT_DIR)/...
