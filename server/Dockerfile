FROM golang:1.24-alpine AS builder

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /server/smart-mirror-backend

# Development stage (target: dev)
FROM golang:1.24-alpine AS dev

WORKDIR /server

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["air"]

# Production stage (target: prod)
FROM alpine:latest AS prod

WORKDIR /server

COPY --from=builder /server/smart-mirror-backend /server/smart-mirror-backend

CMD ["/server/smart-mirror-backend"]