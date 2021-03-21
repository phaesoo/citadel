# Meta
NAME := citadel

# Install dependencies
.PHONY: deps
deps:
	go mod download

# Build the main executable
main:
	go build -o main .

# This is a specialized build for running the executable inside a minimal scratch container
.PHONY: build-docker
build-docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o ./main .

# Run all unit tests
.PHONY: test
test: main
	go test -short ./...

# Run all benchmarks
.PHONY: bench
bench:
	go test -short -bench=. ./...

# test with coverage turned on
.PHONY: cover
cover:
	go test -short -cover -covermode=atomic ./...

# integration test with coverage and the race detector turned on
.PHONY: test-ci
test-ci:
	# go run db/migrate/main.go -t=true
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

# Apply https://golang.org/cmd/gofmt/ to all packages
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: fmt-check
fmt-check:
ifneq ($(shell gofmt -l .),)
	$(error gofmt fail in $(shell gofmt -l .))
endif

# Apply https://github.com/golangci/golangci-lint to changes since forked from main branch
.PHONY: lint
lint:
	golangci-lint run --timeout=5m --new-from-rev=$(shell git merge-base $(shell git branch | sed -n -e 's/^\* \(.*\)/\1/p') origin/main) --enable=unparam --enable=misspell --enable=prealloc

# Remove all compiled binaries from the directory
.PHONY: clean
clean:
	go clean

# Analyze the code for any unused dependencies
.PHONY: prune-deps
prune-deps:
	go mod tidy
