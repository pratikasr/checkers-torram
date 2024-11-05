BINARY_NAME=checkd
MAIN_PATH=./cmd/checkers-torramd/main.go
VERSION := $(shell git describe --tags)
COMMIT := $(shell git log -1 --format='%H')
GOBIN = $(shell go env GOPATH)/bin

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=checkers-torram \
        -X github.com/cosmos/cosmos-sdk/version.AppName=checkers-torram \
        -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
        -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

.PHONY: all build install clean test test-verbose test-coverage

all: install

build:
	@echo "Building Checkers binary..."
	@go build $(BUILD_FLAGS) -o $(GOBIN)/$(BINARY_NAME) $(MAIN_PATH)

install: build
	@echo "Checkers binary installed at: $(GOBIN)/$(BINARY_NAME)"

clean:
	@echo "Removing Checkers binary..."
	@rm -f $(GOBIN)/$(BINARY_NAME)

test:
	@echo "Running tests..."
	@go test ./x/checkerstorram/keeper/...

test-verbose:
	@echo "Running tests with verbose output..."
	@go test -v ./x/checkerstorram/keeper/...

test-coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./x/checkerstorram/keeper/...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"