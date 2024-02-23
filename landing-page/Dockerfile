FROM golang:latest AS build
WORKDIR /go/src
COPY go.mod .
RUN --mount=type=cache,target=/root/.cache/go-build go mod download
COPY handlers handlers/
COPY landing-page.go .
RUN --mount=type=cache,target=/root/.cache/go-build go build -v -o /go/bin/landing-page

FROM ubuntu:latest
COPY --from=build /go/bin/landing-page /bin/landing-page
WORKDIR /bin
ENTRYPOINT ["landing-page"]
