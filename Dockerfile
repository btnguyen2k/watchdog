## Dockerfile to build and package the project as a Docker image
# Sample build command:
# docker build --rm -t watchdog .

FROM golang:1.13-alpine
LABEL maintainer="Thanh Nguyen <btnguyen2k@gmail.com>"
RUN mkdir /build && mkdir /app
COPY . /build
RUN cd /build \
    && go build -o main \
    && cp main /app
WORKDIR /app
EXPOSE 8080
CMD ["/app/main"]
