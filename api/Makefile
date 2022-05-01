.PHONY: build
build: format init
	go build

.PHONY: format
format:
	gofmt -w -d .
	swag fmt

.PHONY: init
init:
	swag init

.PHONY: run
run: format init
	go run main.go

.PHONY: air
air: format init
	air

.PHONY: migrate
migrate:
	go run db/migrate/migrate.go

.PHONY: seed
seed:
	go run db/seeds/seed.go
