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

# Kill process on specific PORT (usage: make kill PORT=3000)
kill:
	@if [ -z "$(PORT)" ]; then \
		echo "Usage: make kill PORT=8080"; \
		exit 1; \
	fi; \
	PIDS=$$(lsof -ti:$(PORT)); \
	if [ -n "$$PIDS" ]; then \
		echo "$$PIDS" | xargs kill -9; \
		echo "Killed PIDs $$PIDS on port $(PORT)"; \
	else \
		echo "No process on port $(PORT)"; \
	fi

# Clean
clean:
	rm -f bin/server coverage.out

# Coverage HTML
cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

# CI
ci: test lint build

