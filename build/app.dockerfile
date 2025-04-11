ARG GO_VER=1.23

FROM golang:${GO_VER} AS builder

WORKDIR /opt

COPY ./cmd/app/ ./cmd/app/
COPY ./internal/ ./internal
COPY go.mod ./go.mod

RUN go mod download && \
    go mod tidy

RUN go build -o ./bin/app ./cmd/app

WORKDIR /opt

FROM scratch

COPY --from=builder /opt/bin/app /app

