PROJECT_NAME=relay

.DEFAULT_GOAL := test

build: clean
	goreleaser build

build-snapshot: clean
	goreleaser build --snapshot

bun-new-migration: export BUN_TIMESTAMP=$(shell date +%Y%m%d%H%M%S | head -c 14)
bun-new-migration:
	touch internal/db/bun/migrations/${BUN_TIMESTAMP}_new.go
	cat internal/db/bun/migrations/migration.go.tmpl > internal/db/bun/migrations/${BUN_TIMESTAMP}_new.go

clean:
	@echo cleaning up workspace
	@rm -Rvf coverage.txt dist relay
	@find . -name ".DS_Store" -exec rm -v {} \;

docker-restart: docker-stop docker-start

docker-start:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml up -d

docker-stop:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml down

fmt:
	@echo formatting
	@go fmt $(shell go list ./... | grep -v /vendor/)

i18n-extract:
	goi18n extract -outdir locales

i18n-merge:
	goi18n merge -outdir locales locales/active.*.toml locales/translate.*.toml

i18n-translations:
	goi18n merge -outdir locales locales/active.*.toml

lint:
	@echo linting
	@golint $(shell go list ./... | grep -v /vendor/)

test: tidy fmt lint #gosec
	go test -cover ./...

test-ext: tidy fmt lint #gosec
	go test --tags=postgres -cover ./...

test-bench-ext: tidy fmt lint #gosec
	go test  -run=XXX -bench=. --tags=postgres -cover ./...

tidy:
	go mod tidy

vendor: tidy
	go mod vendor

.PHONY: build-snapshot bun-new-migration clean fmt lint stage-static npm-scss npm-upgrade docker-restart docker-start docker-stop test test-ext test-race test-race-ext test-verbose tidy vendor
