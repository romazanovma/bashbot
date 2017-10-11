FROM golang:1.7

COPY ./vendors/ /go/src/
COPY ./src/ /go/src/

WORKDIR /go

RUN go build src/bashbot/bashbot.go
