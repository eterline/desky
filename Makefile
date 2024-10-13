.PHONY: build

build:
	go build -v ./cmd/desky/...

.DEFAULT_GOAL := build