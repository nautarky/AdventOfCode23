.DEFAULT_GOAL := run

sources := advent.go helpers.go

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build $(sources)
.PHONY:build

run: vet
	go run $(sources)
.PHONY:run
