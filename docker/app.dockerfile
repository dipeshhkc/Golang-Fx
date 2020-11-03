FROM golang:alpine

WORKDIR /golang-fx

ADD . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./golang-fx"
