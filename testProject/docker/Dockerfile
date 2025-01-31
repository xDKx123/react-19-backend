FROM golang:1.23.4-alpine3.21 AS builder
#FROM cosmtrek/air
LABEL authors="kovacd"

WORKDIR /app

# RUN apk add --no-cache git gcc g++ musl-dev

COPY go.mod go.sum main.go ./
RUN go mod download

# Install air
RUN go install github.com/air-verse/air@latest && air -v
# Install debugger
RUN go install github.com/go-delve/delve/cmd/dlv@latest && dlv version

COPY src/ ./src
COPY static/ ./static
COPY vendor/ ./vendor
COPY main.go ./
COPY .air.toml ./

RUN mkdir -p /app/tmp

#RUN go build -o main ./main.go
# RUN go build -gcflags="all=-N -l" -o compiledApp

#FROM debian:buster
#FROM gcr.io/distroless/static:nonroot
#LABEL authors="kovacd"

EXPOSE 8080 40000

RUN ls -lah /app

# RUN air bench
#WORKDIR /
#COPY --from=builder /go/bin/dlv /
#COPY --from=builder /app/compiledApp /

# CMD ["./main"]
#CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./compiledApp"]

# CMD ["air", "-c", ".air.toml"]

#CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "tmp/main", "&", "air", "-c", ".air.toml"]
CMD ["sh", "-c", "dlv --listen=:40000 --headless=true --api-version=2 --accept-multiclient exec tmp/main & air -c .air.toml"]
