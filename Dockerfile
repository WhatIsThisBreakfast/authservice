FROM golang as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN make

FROM debian:stable-slim

COPY --from=builder /app/server /usr/local/bin/
COPY --from=builder /app/configs/config.toml /usr/local/.config/

ENTRYPOINT [ "/usr/local/bin/server", "-config", "/usr/local/.config/config.toml" ]
