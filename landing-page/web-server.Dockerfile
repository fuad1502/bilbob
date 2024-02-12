FROM golang:latest

WORKDIR /go/src
COPY go.mod .
RUN go mod download && go mod verify

COPY web-server.go .
RUN go build -o /go/bin/web-server

WORKDIR /go/bin
