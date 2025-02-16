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
	docker compose -f deployments/docker/name-docker-compose.yml build

up:
	docker compose -f deployments/docker/name-docker-compose.yml up -d

stop:
	docker compose -f deployments/docker/name-docker-compose.yml down

run: lint test build up
	@echo "Сервис успешно запущен!"

swagger:
	swag init -g cmd/name/main.go -o api/docs

debug:
	DEBUG=true
	docker compose -f deployments/docker/debug-docker-compose.yml up -d

debug-stop:
	docker compose -f deployments/docker/debug-docker-compose.yml down
	DEBUG=false