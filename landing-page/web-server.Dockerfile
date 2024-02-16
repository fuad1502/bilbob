FROM golang:latest AS base
WORKDIR /go/src
COPY go.mod .
RUN --mount=type=cache,target=/root/.cache/go-build go mod download
COPY web-server.go .
RUN --mount=type=cache,target=/root/.cache/go-build go build -v -o /go/bin/web-server

FROM ubuntu:latest
COPY --from=base /go/bin/web-server /bin/web-server
WORKDIR /go/bin
ENTRYPOINT ["web-server"]
