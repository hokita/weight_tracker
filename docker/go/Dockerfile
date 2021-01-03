FROM golang:1.15.6-alpine
# For Raspberry Pi
# FROM arm32v6/golang:1.15.6-alpine

ENV GOBIN /go/bin
ENV GO111MODULE=on
ENV GOPATH=

WORKDIR /go
ADD ./app /go

RUN go build -o weight_tracker main.go
