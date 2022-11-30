package converter

import (
	"example/grpc/internal/controller/rpcg/protobuf"
	"example/grpc/internal/core/entity"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertBook(book entity.Book) *protobuf.Book {
	return &protobuf.Book{
		Id:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		PublishedAt: timestamppb.New(*book.PublishedAt),
		CreatedAt:   timestamppb.New(*book.CreatedAt),
		UpdatedAt:   timestamppb.New(*book.UpdatedAt),
	}
}

func ConvertListBook(books []entity.Book) []*protobuf.Book {
	result := make([]*protobuf.Book, len(books), cap(books))
	for _, book := range books {
		result = append(result, ConvertBook(book))
	}
	return result
}
