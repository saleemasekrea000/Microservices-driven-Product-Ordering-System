
include .env
VERBOSE := 1
## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]





## run: run the cmd/api application
.PHONY: run
run:
	go run ./cmd


## build: build the cmd/api application
.PHONY: build
build:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api

## migration name=$1: create a new database migration
.PHONY: migration
migration:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}
## up: apply all up database migrations
.PHONY: up
up:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${dns} up


## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...


## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

.PHONY: test
test:
.PHONY: test
test:
	docker run --name order-db -e POSTGRES_PASSWORD=1 -p 5501:5432 -d postgres
	
	 while ! pg_isready -h 0.0.0.0 -p 5501 -U postgres; do sleep 1; done

	 docker exec order-db psql -U postgres -c "CREATE USER tester WITH PASSWORD '1' SUPERUSER;"
   
	 docker exec order-db psql -U postgres -c 'CREATE DATABASE test;'


	@echo 'Running down migrations for test database...'
	migrate -path ./migrations -database ${test_dns} down

	@echo 'Running up migrations for test database...'
	migrate -path ./migrations -database "${test_dns}" up

	@echo 'Running tests...'
	go test -vet=off ./...
	docker stop order-db
	docker remove order-db