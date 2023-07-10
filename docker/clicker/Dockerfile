FROM golang:1.20 AS build

COPY . /src

WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -mod mod -a -installsuffix cgo -o /bin/clicker ./cmd/clicker/main.go

RUN mkdir /app &&\
    cp /bin/clicker /app/clicker &&\
    cp /src/internal/config/config.yaml /app/config.yaml

ENV CONFIG_FILE="/app/config.yaml"
ENV HTTP_PROXY="socks5://torproxy:9050"
ENV HTTPS_PROXY="socks5://torproxy:9050"

WORKDIR /app

ENTRYPOINT ["./clicker"]
