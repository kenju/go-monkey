NAME := monkey
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
	-X 'main.revision=$(REVISION)'

packages := ast lexer main.go parser repl token

## Build binaries and run
run: build
	./go-monkey

## Setup
setup:
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

## Install dependencies
deps:
	echo "Not implemented yet"

## Update dependencies
update:
	echo "Not implemented yet"

## Run tests
test: setup
	go test -v ./...

## Lint
lint: setup
	go vet $$(go list)
	for pkg in $$(go list); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

## Format source codes
fmt:
	for pkg in $(packages); do \
		goimports -w $$pkg; \
	done

## Build binaries
build:
	go build -ldflags "$(LDFLAGS)"

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup deps update test lint help
