FROM golang:latest AS build
WORKDIR /go/src
COPY go.mod go.sum .
RUN --mount=type=cache,target=/root/.cache/go-build go mod download
COPY api.go .
COPY dbs/ dbs/
COPY environ/ environ/
COPY errors/ errors/
COPY middlewares/ middlewares/
COPY passwords/ passwords/
COPY routes/ routes/
COPY sessions/ sessions/
RUN --mount=type=cache,target=/root/.cache/go-build go build -v -o /go/bin/api

FROM ubuntu:latest
COPY --from=build /go/bin/api /bin/api
WORKDIR /go/bin
ENTRYPOINT ["api"]
