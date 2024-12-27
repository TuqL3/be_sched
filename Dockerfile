FROM golang:1.23 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y git && \
    go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

EXPOSE 8081

CMD ["air"]
