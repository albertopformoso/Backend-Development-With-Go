# Build stage
FROM golang:1.18.2-alpine3.16 AS builder

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download && \
    go mod verify   && \
    go mod tidy -v

RUN go build -o ./bin/go-api ./cmd/...

# Run stage
FROM alpine:3.16.0

ARG USER="go-api"

RUN adduser \
        -g "log daemon user" \
        --disabled-password \
        ${USER}

WORKDIR /app
RUN mkdir -p cmd/certificates

COPY --from=builder /app/cmd/certificates /app/cmd/certificates
COPY --from=builder /app/bin/go-api /app/bin/go-api

EXPOSE 8080

USER ${USER}:${USER}

ENTRYPOINT [ "/app/bin/go-api" ]