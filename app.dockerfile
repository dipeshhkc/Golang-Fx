FROM golang:alpine

WORKDIR /teamplace

ADD . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./teamplace-api"
