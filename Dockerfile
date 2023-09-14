FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN make build-linux


FROM ubuntu:latest

RUN apt update && apt install ca-certificates -y && rm -rf /var/lib/apt/lists/*
WORKDIR /app

COPY --from=builder /app/bin/pokego /usr/local/bin
