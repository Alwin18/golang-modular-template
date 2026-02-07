APP_NAME=api

run:
	go run cmd/api/main.go

build:
	go build -o bin/$(APP_NAME) cmd/api/main.go

test:
	go test ./... -v

tidy:
	go mod tidy

.PHONY: run build test tidy lint