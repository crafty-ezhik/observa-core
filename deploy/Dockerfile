# Stage 1 — Build
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ENV APP_ENV=docker

COPY ../.. ./
RUN CGO_ENABLED=0 GOOS=linux go build -o eventhub ./cmd/observacore/main.go

# Stage 2 — Final Image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/eventhub .
# COPY ./configs ./configs

# Опционально: если приложение пишет в файл
VOLUME ["./logs"]

EXPOSE 8080

CMD ["./eventhub"]
