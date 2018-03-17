FROM golang:1.10-alpine

RUN apk update && \ 
    apk add curl git && \
    curl https://glide.sh/get | sh && \
    go get github.com/pilu/fresh

ADD glide.yaml /go/src/github.com/glide.yaml
WORKDIR /go/src/github.com

RUN glide install