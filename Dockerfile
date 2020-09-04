## Dockerfile to build and package the project as a Docker image
# Sample build command:
# docker build --rm -t watchdog .

FROM golang:1.13-alpine AS builder
LABEL maintainer="Thanh Nguyen <btnguyen2k@gmail.com>"
RUN mkdir -p /build
COPY . /build
RUN cd /build \
    && go build -o main

FROM alpine:3.10
LABEL maintainer="Thanh Nguyen <btnguyen2k@gmail.com>"
RUN mkdir -p /app
COPY --from=builder /build/main /app/main
RUN apk add --no-cache -U tzdata bash ca-certificates \
    && update-ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime \
    && chmod 711 /app/main \
    && rm -rf /var/cache/apk/*
WORKDIR /app
EXPOSE 8080
CMD ["/app/main"]
