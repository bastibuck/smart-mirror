FROM golang:1.24-alpine

WORKDIR /server

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air"]