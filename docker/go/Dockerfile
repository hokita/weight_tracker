FROM golang:1.16.0-alpine

ENV GOPATH=

WORKDIR /go
ADD ./app /go

RUN go build -o weight_tracker main.go
