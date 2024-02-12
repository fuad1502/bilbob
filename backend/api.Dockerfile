FROM golang:latest

WORKDIR /go/src
COPY go.mod go.sum .
RUN go mod download && go mod verify

COPY api.go .
RUN go build -o /go/bin/api

WORKDIR /go/bin
