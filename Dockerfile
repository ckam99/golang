FROM golang:1.18-alpine

RUN apk add --no-cache curl git

WORKDIR /app


RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air



