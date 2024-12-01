SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
.SILENT:

# go builds are fast enough that we can just build on demand instead of trying to do any fancy
# change detection
build: clean advent
.PHONY: build

advent:
	go build ./cmd/advent

clean:
	rm -f ./advent
.PHONY: clean

test:
	go test ./...
.PHONY: test
