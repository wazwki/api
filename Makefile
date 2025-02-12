.PHONY: lint test build up

DEBUG ?= false

export DEBUG

lint:
	golangci-lint run
	@echo "Линтер пройден"

test:
	go test ./...
	@echo "Тесты пройдены"

build:
	docker compose build

up:
	docker compose up -d

run: lint test build up
	@echo "Сервис успешно запущен!"

swagger:
	swag init -g cmd/main.go -o api/docs