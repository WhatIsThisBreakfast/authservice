.PHONY: build

build:
		go build ./cmd/server

.DEFAULT := build

