FROM golang:latest
WORKDIR /app
COPY . .
RUN make build-linux


# FROM debian:latest

# RUN useradd -c "pokego user" pokego && mkdir /pokegohome && chown pokego:pokego /pokegohome
RUN apt update && apt install ca-certificates -y && rm -rf /var/lib/apt/lists/*
# WORKDIR /app

# USER pokego

COPY ./bin/pokego /usr/local/bin
