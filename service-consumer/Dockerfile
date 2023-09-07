FROM golang:1.21-alpine

# Required because go requires gcc to build
RUN apk add build-base

RUN apk add inotify-tools

COPY . /consumer


RUN apk add bash

ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /consumer

USER root

RUN go mod download -x -json
RUN go mod vendor

CMD go build /consumer/ && /consumer/consumer
