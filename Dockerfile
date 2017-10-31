FROM golang:1.7

COPY ./vendors/ /go/src/
COPY ./src/ /go/src/


RUN go build /go/src/bashbot/bashbot.go
