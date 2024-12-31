FROM golang:1.23.4
LABEL authors="kovacd"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ ./src

EXPOSE 8080

RUN go build -o main ./src/main.go

CMD ["./main"]