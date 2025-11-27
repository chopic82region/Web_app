FROM golang:1.25.3

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o main .

# Экспортируем порт
EXPOSE 8080

# Команда запуска
CMD ["./main"]