.DEFAULT_GOAL := run

sources := main.go d1.go d2.go d3.go d4.go d5.go d6.go

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
