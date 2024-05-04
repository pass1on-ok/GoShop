# Используем образ Golang
FROM golang:latest

# Установка PostgreSQL клиента
RUN apt-get update && apt-get install -y postgresql-client

# Создание директории внутри контейнера
RUN mkdir -p /app

# Установка рабочей директории
WORKDIR /app

# Копирование исходного кода внутрь контейнера
COPY . .

# Сборка приложения
RUN go build -o main cmd/main.go

# Команда для запуска вашего приложения
CMD ["./main"]
