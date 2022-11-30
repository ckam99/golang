package converter

import (
	"example/grpc/internal/controller/rpc/protobuf"
	"example/grpc/internal/core/entity"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertBook(book entity.Book) *protobuf.Book {
	b := &protobuf.Book{
		Id:          book.ID,
		Title:       book.Title,
		Description: book.Description,
	}
	if book.CreatedAt != nil {
		b.CreatedAt = timestamppb.New(*book.CreatedAt)
	}
	if book.AuthorID != nil {
		b.AuthorId = *book.AuthorID
	}
	if book.Author != nil {
		b.Author = ConvertAuthor(*book.Author)
	}
	if book.PublishedAt != nil {
		b.PublishedAt = timestamppb.New(*book.PublishedAt)
	}
	if book.UpdatedAt != nil {
		b.UpdatedAt = timestamppb.New(*book.UpdatedAt)
	}
	return b
}
