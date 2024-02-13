FROM golang:latest

WORKDIR /go/src
COPY go.mod go.sum .
RUN go mod download && go mod verify

COPY api.go .
COPY passwords/passwords.go passwords/
COPY errors/errors.go errors/
COPY dbs/dbs.go dbs/
COPY middlewares/middlewares.go middlewares/
COPY routes/routes.go routes/
RUN go build -o /go/bin/api

WORKDIR /go/bin
