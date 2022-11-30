FROM golang:1.19-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache curl git make gcc protobuf protobuf-dev 

# FROM golang:1.19
# RUN apt-get update -y && apt-get install -y curl make libprotobuf-dev protobuf-dev protobuf-compiler

WORKDIR /app

# install go-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
RUN mv migrate /usr/local/bin/migrate

# install go protobuf packages
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28  && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# instal dev packages
RUN go install -v github.com/cweill/gotests/gotests@v1.6.0 && \
    go install -v github.com/fatih/gomodifytags@v1.16.0 && \
    go install -v github.com/josharian/impl@v1.1.0 && \
    go install -v github.com/haya14busa/goplay/cmd/goplay@v1.0.0 && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install -v honnef.co/go/tools/cmd/staticcheck@latest && \
    go install -v golang.org/x/tools/gopls@latest



