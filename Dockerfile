# Stage 1: Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ code và build
COPY . .
RUN go build -o main cmd/main.go

# Stage 2: Run stage (Production Image)
FROM alpine:latest

RUN apk add --no-cache curl

WORKDIR /app

# Chỉ copy file binary từ stage builder sang, giúp image nhẹ đi ~300MB
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]