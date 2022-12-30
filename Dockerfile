FROM golang:1.19-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache curl git make

WORKDIR /app