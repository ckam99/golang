package rpcg

import (
	"context"
	"example/grpc/internal/controller/rpcg/converter"
	"example/grpc/internal/controller/rpcg/protobuf"
	"example/grpc/internal/core/entity"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetBooks(rq *protobuf.QueryRequest, stream protobuf.BookService_GetBooksServer) error {
	books, err := s.bookService.GetAll(context.Background())
	if err != nil {
		return status.Errorf(codes.Internal, "failed to fetch books: %s", err)
	}
	for _, book := range books {
		if err := stream.Send(converter.ConvertBook(book)); err != nil {
			return status.Errorf(codes.Internal, "failed to stream books: %s", err)
		}
	}
	return nil
}

func (s *Server) CreateBook(ctx context.Context, in *protobuf.CreateBookRequest) (*protobuf.Book, error) {
	dt := in.GetPublishedAt().AsTime()
	book := entity.Book{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
		PublishedAt: &dt,
	}
	if err := s.bookService.Create(ctx, &book); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to persist book:%s", err)
	}
	return converter.ConvertBook(book), nil
}

func (s *Server) UpdateBook(ctx context.Context, in *protobuf.UpdateBookRequest) (*protobuf.Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "not implemented")
}
