.PHONY: build test clean run install

# Version info
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Build flags
LDFLAGS := -ldflags "-X github.com/KilimcininKorOglu/gesh/pkg/version.Version=$(VERSION) \
                     -X github.com/KilimcininKorOglu/gesh/pkg/version.Commit=$(COMMIT) \
                     -X github.com/KilimcininKorOglu/gesh/pkg/version.BuildDate=$(BUILD_DATE)"

# Default target
all: build

# Build the binary
build:
	go build $(LDFLAGS) -o gesh .

# Build for all platforms
build-all:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/gesh-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/gesh-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/gesh-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/gesh-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/gesh-windows-amd64.exe .

# Run tests
test:
	go test ./... -v

# Run tests with coverage
test-cover:
	go test ./... -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -f gesh gesh.exe
	rm -rf dist/
	rm -f coverage.out coverage.html

# Run the editor
run: build
	./gesh

# Install to GOPATH/bin
install:
	go install $(LDFLAGS) .

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Show version
version:
	@echo "Version: $(VERSION)"
	@echo "Commit: $(COMMIT)"
	@echo "Build Date: $(BUILD_DATE)"
