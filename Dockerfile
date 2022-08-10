FROM golang:1.18-alpine

RUN apk add --no-cache curl git make

WORKDIR /app


RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air


RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

RUN mv migrate /usr/local/bin/migrate
