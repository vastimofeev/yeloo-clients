SHELL := /bin/bash

# Команды
.PHONY: all build run test docker-build docker-up docker-down swag-clean swag-init

# Сборка бинарного файла
build:
	go build -o main cmd/server/main.go

# Запуск приложения локально
run:
	go run cmd/server/main.go

# Запуск тестов
test:
	go test ./... -v

# Генерация Swagger-документации
swag-init:
	swag init -g cmd/server/main.go

# Очистка Swagger-документации
swag-clean:
	rm -rf docs/*

# Сборка Docker-образа
docker-build:
	docker-compose build

# Запуск Docker-контейнеров
docker-up:
	docker-compose up --build

# Остановка Docker-контейнеров
docker-down:
	docker-compose down

# Полная очистка Docker (включая тома)
docker-clean:
	docker-compose down -v

# Запуск приложения в Docker
docker-run:
	docker-compose up --build -d
	@echo "Ожидание состояния healthy для контейнера yeloo-clients-app-1..."
	@while [ "$$(docker inspect -f '{{.State.Health.Status}}' yeloo-clients-app-1)" != "healthy" ]; do \
		echo "Container yeloo-clients-app-1 not ready yet. Waiting..."; \
		sleep 5; \
	done
	@echo "Container yeloo-clients-app-1 healthy!"

# Помощь
help:
	@echo "Доступные команды:"
	@echo "  build			- Сборка бинарного файла"
	@echo "  run			- Запуск приложения локально"
	@echo "  test			- Запуск тестов"
	@echo "  swag-init		- Генерация Swagger-документации"
	@echo "  swag-clean		- Очистка Swagger-документации"
	@echo "  docker-build	- Сборка Docker-образа"
	@echo "  docker-up		- Запуск Docker-контейнеров"
	@echo "  docker-down	- Остановка Docker-контейнеров"
	@echo "  docker-clean	- Полная очистка Docker (включая тома)"
	@echo "  docker-run		- Запуск приложения в Docker"