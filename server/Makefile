# Include variables from the .envrc file
include .envrc


# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api -dsn=${TEMA_DB_DSN}

## db/migrations/version: check current database migration version
.PHONY: db/migrations/version
db/migrations/version:
	goose -dir ./migrations postgres ${TEMA_DB_DSN} version

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating a migration file for ${name}'
	goose -s -dir ./migrations create ${name} sql

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	goose -dir ./migrations postgres ${TEMA_DB_DSN} up

## db/migrations/down: apply a down database migration
.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo 'Running down migrations...'
	goose -dir ./migrations postgres ${TEMA_DB_DSN} down



# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor



# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the cmd/api application for local machine and linux/amd64
.PHONY: build/api
build/api:
	@echo 'Building cmd/api for local machine...'
	go build -ldflags='-s -w' -o=./bin/local/api ./cmd/api
	@echo 'Building cmd/api for deployment in linux/amd64...'
	GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o=./bin/linux_amd64/api ./cmd/api