VERSION 0.6
FROM golang:1.18-bullseye
WORKDIR /papergirl

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

test:
    FROM +deps
    COPY *.go .
    COPY --dir app app
    RUN go test ./...

build:
    FROM +test
    RUN go build -o papergirl app/*
    SAVE ARTIFACT papergirl /papergirl AS LOCAL build/papergirl

docker:
    FROM ubuntu:20.04
    COPY +build/papergirl .
    ENTRYPOINT ["/papergirl"]
    SAVE IMAGE papergirl:latest
