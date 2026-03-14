.PHONY: test lint build run clean all pre-commit

# Pre-commit (MANDATORY before commit)
pre-commit: test lint build

# All existing (test lint build)
all: test lint build

# Run all tests + coverage
test:
	go test ./... -v -race -coverprofile=coverage.out
	go tool cover -func=coverage.out

# Lint
lint:
	golangci-lint run

# Install linter
lint-install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Build
build:
	go build -v -o bin/server ./cmd/server

# Run dev server
run: build
	./bin/server

# Clean
clean:
	rm -f bin/server coverage.out

# Coverage HTML
cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

# CI
ci: test lint build

