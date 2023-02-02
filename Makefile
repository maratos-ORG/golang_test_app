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
LDFLAGS = -a -installsuffix cgo -ldflags "-X main.appName=${APPNAME} -X main.gitTag=${TAG} -X main.gitCommit=${COMMIT} -X main.gitBranch=${BRANCH}"
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

build: clean dep ## Build
	mkdir -p ./bin
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o bin/${APPNAME} ./$(PROJECT_DIR)

docker-build: ## Build docker image
	docker build -t boosterkrd/${APPNAME}:${TAG} .
	docker image prune --force --filter label=stage=intermediate
	docker tag boosterkrd/${APPNAME}:${TAG} boosterkrd/${APPNAME}:latest

docker-push: ## Push docker image to the registry
	docker push boosterkrd/${APPNAME}:${TAG}
	docker push boosterkrd/${APPNAME}:latest
