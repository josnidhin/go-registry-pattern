#
# @author Jose Nidhin
#
VERSION := $(shell git describe --always --long --dirty)
PROJECT_NAME := $(shell basename "$(PWD)")
TEST_NAME := $(PROJECT_NAME)_TEST
GO_SRC_FILES := $(shell find . -type f -name '*.go')
GO_SRC_MAIN := $(shell ls *.go)

CONTAINER_VERSION=latest
GO_ENVFLAGS=CGO_ENABLED=0
LDFLAGS=-ldflags "-X=buildinfo.AppVersion=$(VERSION) -X=buildinfo.AppName=$(PROJECT_NAME)"

all: generate tidy vet fmt simplify test clean build

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	gofmt -l -w $(GO_SRC_FILES)

.PHONY: simplify
simplify:
	gofmt -s -l -w $(GO_SRC_FILES)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: generate
generate:
	go generate ./...

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: clean
clean:
	go clean -x
	rm -f cover.out coverage.html

.PHONY: build
build:
	go build -v -o $(PROJECT_NAME) $(LDFLAGS) $(GO_SRC_MAIN)

.PHONY: coverage
coverage:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o coverage.html

.PHONY: release-build
release-build:
	$(GO_ENVFLAGS) go build -v -o $(PROJECT_NAME) $(LDFLAGS) $(GO_SRC_MAIN)

.PHONY: container-clean-build
container-clean-build:
	docker image build --no-cache -t $(PROJECT_NAME):$(CONTAINER_VERSION) .

.PHONY: container-build
container-build:
	docker image build -t $(PROJECT_NAME):$(CONTAINER_VERSION) .
