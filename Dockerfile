# ---------- Build stage ----------
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY . .

# Собираем бинарник (linux, static)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app

# ---------- Run stage ----------
FROM alpine:3.20

WORKDIR /app

# сертификаты (если есть HTTPS-запросы)
RUN apk add --no-cache ca-certificates

# копируем бинарник из builder
COPY --from=builder /app/app .

# порт (поменяй под свой сервис)
EXPOSE 8080

# запуск
CMD ["./app"]