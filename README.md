# gRPC with buf

# install protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28  && \
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# gRPC gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.14.0 && \
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.14.0

# install evans
go install github.com/ktr0731/evans@latest

