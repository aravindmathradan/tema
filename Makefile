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

## run/api: run the cmd/api application in development mode
.PHONY: run/api
run/api:
	cd server && go run ./cmd/api -dsn=${TEMA_DB_DSN}

## run/frontend: run the app frontend in development mode
.PHONY: run/frontend
run/frontend:
	cd frontend && npm run dev

## db/migrations/version: check current database migration version
.PHONY: db/migrations/version
db/migrations/version:
	goose -dir ./server/migrations postgres ${TEMA_DB_DSN} version

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating a migration file for ${name}'
	goose -s -dir ./server/migrations create ${name} sql

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	goose -dir ./server/migrations postgres ${TEMA_DB_DSN} up

## db/migrations/down: apply a down database migration
.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo 'Running down migrations...'
	goose -dir ./server/migrations postgres ${TEMA_DB_DSN} down


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	cd server && go fmt ./...
	@echo 'Vetting code...'
	cd server && go vet ./...
	cd server && staticcheck ./...
	@echo 'Running tests...'
	cd server && go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	cd server && go mod tidy
	cd server && go mod verify
	@echo 'Vendoring dependencies...'
	cd server && go mod vendor


# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the cmd/api application for local machine and linux/amd64
.PHONY: build/api
build/api:
	@echo 'Building cmd/api for local machine...'
	cd server && go build -ldflags='-s -w' -o=./bin/local/api ./cmd/api
	@echo 'Building cmd/api for deployment in linux/amd64...'
	cd server && GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o=./bin/linux_amd64/api ./cmd/api

## build/frontend: build the sveltekit frontend application for production
.PHONY: build/frontend
build/frontend:
	@echo 'Building frontend for production...'
	cd frontend \
		&& npm run build \
		&& cp package.json package-lock.json build/ \
		&& cd ./build \
		&& npm ci --omit dev


# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #

production_host_ip = '178.128.56.85'

## production/connect: connect to the production server
.PHONY: production/connect
production/connect:
	ssh tema@${production_host_ip}

## production/deploy/api: deploy the api to production
.PHONY: production/deploy/api
production/deploy/api:
	@echo 'Deploying api server on production...'
	rsync -P ./server/bin/linux_amd64/api tema@${production_host_ip}:~
	rsync -rP --delete ./server/migrations tema@${production_host_ip}:~
	rsync -P ./remote/production/api.service tema@${production_host_ip}:~
	rsync -P ./remote/production/Caddyfile tema@${production_host_ip}:~
	ssh -t tema@${production_host_ip} '\
		goose -dir ~/migrations postgres $${TEMA_DB_DSN} up \
		&& sudo mv ~/api.service /etc/systemd/system/ \
		&& sudo systemctl enable api \
		&& sudo systemctl restart api \
		&& sudo mv ~/Caddyfile /etc/caddy/ \
		&& sudo systemctl reload caddy \
	'

## production/deploy/frontend: deploy the sveltekit frontend application to production
.PHONY: production/deploy/frontend
production/deploy/frontend:
	@echo 'Deploying frontend application on production...'
	rsync -rP --delete ./frontend/build/ tema@${production_host_ip}:~/frontend/
	ssh -t tema@${production_host_ip} 'bash -i -c "\
		pm2 start ~/frontend/index.js --name frontend \
	"'