FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/smart-mirror-backend

# Development stage (target: dev)
FROM golang:1.24-alpine AS dev

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["air"]

# Production stage (target: prod)
FROM alpine:latest AS prod

WORKDIR /app

COPY --from=builder /app/smart-mirror-backend /app/smart-mirror-backend

CMD ["/app/smart-mirror-backend"]