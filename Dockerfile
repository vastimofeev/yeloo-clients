FROM golang:1.24.2 AS runner

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Устанавливаем рабочую директорию для запуска
WORKDIR /app/

# Запускаем приложение через go run
CMD ["go", "run", "cmd/server/main.go"]