package handler

// import (
// 	"context"
// 	"example/grpc/internal/controller/rpc/converter"
// 	"example/grpc/internal/controller/rpc/protobuf"
// 	"example/grpc/internal/core/entity"
// 	"example/grpc/internal/core/ports"
// 	"example/grpc/internal/core/service"
// 	"example/grpc/internal/provider/postgres"
// 	"example/grpc/pkg/postgresql"
// 	"example/grpc/pkg/utils"

// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// type BookServer struct {
// 	protobuf.UnimplementedBookServiceServer
// 	service ports.BookService
// }

// func NewBookServer(db postgresql.Client) *BookServer {
// 	return &BookServer{
// 		service: service.NewBookService(
// 			postgres.NewBookRepository(db),
// 		),
// 	}
// }

// func (s *BookServer) GetBooks(rq *protobuf.QueryRequest, stream protobuf.BookService_GetBooksServer) error {
// 	books, err := s.service.GetAll(context.Background())
// 	if err != nil {
// 		return status.Errorf(codes.Internal, "failed to fetch books: %s", err)
// 	}
// 	for _, book := range books {
// 		if err := stream.Send(converter.ConvertBook(book)); err != nil {
// 			return status.Errorf(codes.Internal, "failed to stream books: %s", err)
// 		}
// 	}
// 	return nil
// }

// func (s *BookServer) CreateBook(ctx context.Context, in *protobuf.CreateBookRequest) (*protobuf.Book, error) {
// 	dt := in.GetPublishedAt().AsTime()
// 	book := entity.Book{
// 		Title:       in.GetTitle(),
// 		Description: in.GetDescription(),
// 		PublishedAt: &dt,
// 	}
// 	if err := s.service.Create(ctx, &book); err != nil {
// 		if err == utils.ErrInvalidForeinKey {
// 			return nil, status.Errorf(codes.InvalidArgument, "author_id does not exists :%s", err)
// 		}
// 		return nil, status.Errorf(codes.Internal, "failed to persist book:%s", err)
// 	}
// 	return converter.ConvertBook(book), nil
// }

// func (s *BookServer) UpdateBook(ctx context.Context, in *protobuf.UpdateBookRequest) (*protobuf.Book, error) {
// 	book := entity.Book{
// 		ID:          in.GetId(),
// 		Description: in.GetDescription(),
// 	}
// 	if err := s.service.Update(ctx, &book); err != nil {
// 		if err == utils.ErrNoEntity {
// 			return nil, status.Errorf(codes.NotFound, "failed to get book :%s", err)
// 		}
// 		return nil, status.Errorf(codes.Internal, "failed to persist book:%s", err)
// 	}
// 	return converter.ConvertBook(book), nil
// }

// func (s *BookServer) DeleteBook(ctx context.Context, in *protobuf.Id) (*protobuf.Empty, error) {
// 	if err := s.service.Delete(ctx, in.GetValue()); err != nil {
// 		if err == utils.ErrNoEntity {
// 			return nil, status.Errorf(codes.NotFound, "failed to get book :%s", err)
// 		}
// 		return nil, status.Errorf(codes.Internal, "failed to delete book:%s", err)
// 	}
// 	return nil, nil
// }

// func (s *BookServer) GetBook(ctx context.Context, in *protobuf.Id) (*protobuf.Book, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "not implemented")
// }
