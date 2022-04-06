FROM golang:latest as build

WORKDIR /project

COPY go.mod .

RUN go mod download

COPY . .

RUN go build ./cmd/server

FROM debian:stable-slim

COPY --from=build /project/server /usr/local/bin/
COPY --from=build /project/configs/server.toml /usr/local/сonfig/server/

USER 1000

ENTRYPOINT ["server", "-config", "/usr/local/сonfig/server/server.toml"]