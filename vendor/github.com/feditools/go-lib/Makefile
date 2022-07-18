PROJECT_NAME=login

.DEFAULT_GOAL := test

check:
	golangci-lint run

check-fix:
	golangci-lint run --fix

fmt:
	@echo formatting
	@go fmt $(shell go list ./... | grep -v /vendor/)

i18n-extract:
	goi18n extract -format yaml -outdir language/locales

i18n-merge:
	goi18n merge -format yaml -outdir language/locales language/locales/active.*.toml language/locales/translate.*.toml

i18n-translations:
	goi18n merge -format yaml -outdir language/locales language/locales/active.*.toml

test: tidy fmt lint
	go test -cover ./...

tidy:
	go mod tidy

.PHONY: check check-fix fmt i18n-extract i18n-merge i18n-translations lint test tidy
