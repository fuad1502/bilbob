FROM node:latest AS build
WORKDIR webapp
COPY webapp/package-lock.json .
COPY webapp/package.json .
RUN --mount=type=cache,target=/root/.cache/node-build npm install
COPY .env ./
COPY webapp/public/ ./public/
COPY webapp/src/ ./src/
RUN --mount=type=cache,target=/root/.cache/node-build npm run build

FROM node:latest AS prod
RUN npm install -g serve
COPY --from=build /webapp/build/ ./build/
ENTRYPOINT ["serve", "-s", "build"]
