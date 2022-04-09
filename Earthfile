VERSION 0.6
FROM golang:1.18-bullseye
WORKDIR /papergirl

deps-app:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

test-app:
    FROM +deps-app
    COPY *.go .
    COPY --dir app app
    RUN go test ./...

build-app:
    FROM +test-app
    RUN go build -o papergirl app/*
    SAVE ARTIFACT papergirl /papergirl AS LOCAL build/papergirl

docker:
    FROM ubuntu:20.04
    COPY +build-app/papergirl .
    ENTRYPOINT ["/papergirl"]
    SAVE IMAGE papergirl:latest
