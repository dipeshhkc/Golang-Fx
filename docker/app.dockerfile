FROM golang:alpine

WORKDIR /golang-fx

ADD . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ARG VERSION="4.13.0"

RUN set -x \
    && apk add --no-cache git \
    && git clone --branch "v${VERSION}" --depth 1 --single-branch https://github.com/golang-migrate/migrate /tmp/go-migrate

WORKDIR /tmp/go-migrate

RUN set -x \
    && CGO_ENABLED=0 go build -tags 'mysql' -ldflags="-s -w" -o ./migrate ./cmd/migrate \
    && ./migrate -version

RUN cp /tmp/go-migrate/migrate /usr/bin/migrate

WORKDIR /golang-fx

ENTRYPOINT CompileDaemon -command="./golang-fx"