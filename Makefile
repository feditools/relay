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
	@rm -Rvf coverage.txt dist gosec.xml feditools
	@find . -name ".DS_Store" -exec rm -v {} \;
	@rm -Rvf web/bootstrap/dist
	@rm -Rvf web/static/css/bootstrap.min.css web/static/css/bootstrap.min.css.map web/static/css/default.min.css web/static/css/error.min.css web/static/css/login.min.css
	@rm -Rvf web/static/js/bootstrap.bundle.min.js web/static/js/bootstrap.bundle.min.js.map

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

npm-scss:
	cd web/bootstrap && npm run sass

npm-upgrade:
	cd web/bootstrap && npm upgrade

stage-static:
	minify web/static-src/css/default.css > web/static/css/default.min.css
	minify web/static-src/css/error.css > web/static/css/error.min.css
	minify web/static-src/css/login.css > web/static/css/login.min.css
	minify web/bootstrap/dist/bootstrap.css > web/static/css/bootstrap.min.css
	cat web/bootstrap/dist/bootstrap.css.map > web/static/css/bootstrap.min.css.map
	cat web/bootstrap/node_modules/bootstrap/dist/js/bootstrap.bundle.min.js > web/static/js/bootstrap.bundle.min.js
	cat web/bootstrap/node_modules/bootstrap/dist/js/bootstrap.bundle.min.js.map > web/static/js/bootstrap.bundle.min.js.map

test-docker-restart: test-docker-stop test-docker-start

test-docker-start:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml up -d

test-docker-stop:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml down

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

.PHONY: build-snapshot bun-new-migration clean fmt lint stage-static npm-scss npm-upgrade test-docker-restart test-docker-start test-docker-stop test test-ext test-race test-race-ext test-verbose tidy vendor
