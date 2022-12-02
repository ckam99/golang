package interceptor

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

func HttpLogInterceptor(http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func GrpcLogger(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("received grpc request")
	result, err := handler(ctx, req)
	return result, err
}
