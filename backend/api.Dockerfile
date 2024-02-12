FROM golang:latest

WORKDIR /go/src
COPY go.mod go.sum .
RUN go mod download && go mod verify

COPY api.go db.go gin-middlewares.go users-route-handlers.go .
COPY password/password.go password/
RUN go build -o /go/bin/api

WORKDIR /go/bin
