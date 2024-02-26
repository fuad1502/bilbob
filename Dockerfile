FROM golang:latest AS backend-mod-download
WORKDIR /go/src/backend
COPY backend/go.mod .
RUN --mount=type=cache,target=/root/.cache/go-build go mod download

FROM backend-mod-download AS lp-mod-download
WORKDIR /go/src/landing-page
COPY landing-page/go.mod .
RUN --mount=type=cache,target=/root/.cache/go-build go mod download

FROM lp-mod-download AS staticserver-mod-download
WORKDIR /go/src/staticserver
COPY staticserver/go.mod .
RUN --mount=type=cache,target=/root/.cache/go-build go mod download

FROM staticserver-mod-download AS build-staticserver
WORKDIR /go/src
COPY go.work .
COPY go.work.sum .
COPY backend/ backend/
COPY landing-page/ landing-page/
COPY staticserver/ staticserver/
WORKDIR /go/src/staticserver
RUN --mount=type=cache,target=/root/.cache/go-build go build -v -o /go/bin/staticserver

FROM node:latest AS build-webapp
WORKDIR /node/webapp
COPY webapp/package-lock.json .
COPY webapp/package.json .
RUN --mount=type=cache,target=/root/.cache/node-build npm install
COPY .env .
RUN sed -i -e 's/^/REACT_APP_/' .env
COPY webapp/public/ public/
COPY webapp/src/ src/
RUN --mount=type=cache,target=/root/.cache/node-build npm run build

FROM ubuntu:latest
COPY --from=build-staticserver /go/bin/staticserver /bin/staticserver
COPY --from=build-webapp /node/webapp/build/ /bin/build/
WORKDIR /go/bin
ENTRYPOINT ["staticserver"]
