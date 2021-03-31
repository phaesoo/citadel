# Meta
NAME := keybox

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

# Build proto files
.PHONY: proto
proto:
	protoc -I . --go_out ./gen/go \
	--go_opt paths=source_relative \
	--go-grpc_out ./gen/go \
	--go-grpc_opt paths=source_relative \
	proto/*.proto

# Build proto gateway
.PHONY: proto-gw
proto-gw:
	protoc -I . --grpc-gateway_out ./gen/gw \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt grpc_api_configuration=./api_config.yaml \
	--grpc-gateway_opt standalone=true \
	proto/*.proto

# Migrate rdb.
.PHONY: migrate
migrate-mysql:
	go run db/migrate/main.go

# Create a new empty migration file.
.PHONY: migration
migration:
	$(eval VER := $(shell date +"%Y%m%d%H%M%S"))
	$(eval FILE := db/migrate/migrations/migration_$(VER).go)
	cp db/migrate/migrations/template.txt $(FILE)
	sed -i "s/MIGRATION_ID/$(VER)/g" $(FILE)
