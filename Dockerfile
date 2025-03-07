# Используем официальный образ Golang для сборки
FROM golang:1.23.2 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем код проекта
COPY . .

# Загружаем зависимости и собираем бинарник
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/app cmd/main.go

# Финальный образ на базе Alpine (легковесный)
FROM alpine:latest
WORKDIR /app/

# Копируем собранный бинарник
COPY --from=builder /app/app /app/app

# Копируем конфигурационные файлы
COPY config.json ./
COPY .env ./

# Создаём нужные директории
RUN mkdir -p logs data

# Делаем бинарник исполняемым
RUN chmod +x /app/app

# Запускаем приложение
CMD ["/app/app"]
