FROM golang:1.14-alpine

ENV GO111MODULE=on

WORKDIR /go/src/github.com/kons16/team7-backend

COPY . .
RUN apk add --no-cache && go get github.com/pilu/fresh
