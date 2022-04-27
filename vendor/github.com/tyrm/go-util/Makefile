fmt:
	go fmt ./...

lint:
	@echo linting
	@golint $(shell go list ./... | grep -v /vendor/)

test: tidy fmt lint
	go test -cover ./...

test-race: tidy fmt lint
	go test -race -cover ./...

test-verbose: tidy fmt lint
	go test -v -cover ./...

tidy:
	go mod tidy

.PHONY: fmt lint test