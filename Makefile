# Переменные
APP_NAME = yeloo-clients
DOCKER_COMPOSE = docker-compose
GO = go
SWAG = swag
SERVER_PORT = 8080

# Команды
.PHONY: all build run test docker-build docker-up docker-down swag-clean swag-init

# Сборка бинарного файла
build:
	$(GO) build -o main cmd/server/main.go

# Запуск приложения локально
run:
	$(GO) run cmd/server/main.go

# Запуск тестов
test:
	$(GO) test ./... -v

# Генерация Swagger-документации
swag-init:
	$(SWAG) init -g cmd/server/main.go

# Очистка Swagger-документации
swag-clean:
	rm -rf docs/*

# Сборка Docker-образа
docker-build:
	$(DOCKER_COMPOSE) build

# Запуск Docker-контейнеров
docker-up:
	$(DOCKER_COMPOSE) up --build

# Остановка Docker-контейнеров
docker-down:
	$(DOCKER_COMPOSE) down

# Полная очистка Docker (включая тома)
docker-clean:
	$(DOCKER_COMPOSE) down -v

# Запуск приложения в Docker
docker-run:
	$(DOCKER_COMPOSE) up --build -d

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