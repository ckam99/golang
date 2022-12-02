package handler

import (
	"context"
	"example/grpc/internal/controller/rpc/converter"
	"example/grpc/internal/controller/rpc/pb"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/ports"
	"example/grpc/internal/core/service"
	"example/grpc/internal/provider/postgres"
	"example/grpc/pkg/postgresql"
	"example/grpc/pkg/utils"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookServer struct {
	pb.UnimplementedBookServiceServer
	service ports.BookService
	logger  *log.Logger
}

func NewBookServer(db postgresql.Client, logger *log.Logger) *BookServer {
	return &BookServer{
		service: service.NewBookService(
			postgres.NewBookRepository(db),
		),
		logger: logger,
	}
}

func (s *BookServer) StreamBooks(q *pb.Empty, stream pb.BookService_StreamBooksServer) error {
	books, err := s.service.GetAll(context.Background(), 0, 0)
	if err != nil {
		s.logger.Println(err)
		return status.Errorf(codes.Internal, "failed to fetch books: %s", err)
	}
	for _, book := range books {
		if err := stream.Send(converter.ConvertBook(book)); err != nil {
			s.logger.Println(err)
			return status.Errorf(codes.Internal, "failed to stream books: %s", err)
		}
	}
	return nil
}

func (s *BookServer) GetBooks(ctx context.Context, q *pb.QueryRequest) (*pb.BookListResponse, error) {
	count, err := s.service.Count(ctx, 0, 0)
	if err != nil {
		s.logger.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to get authors count: %s", err)
	}
	var pageCount int64
	if q.Limit > 0 {
		pageCount = int64(math.Ceil(float64(count) / float64(q.Limit)))
	}
	offset := q.Offset*q.Limit - q.Limit
	books, err := s.service.GetAll(ctx, q.Limit, offset)
	if err != nil {
		s.logger.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to fetch books: %s", err)
	}
	return &pb.BookListResponse{
		PageCount: pageCount,
		Page:      q.Offset,
		Total:     count,
		Books:     converter.ConvertListBook(books)}, nil
}

func (s *BookServer) CreateBook(ctx context.Context, in *pb.CreateBookRequest) (*pb.Book, error) {
	dt := in.GetPublishedAt().AsTime()
	authorID := in.GetAuthorId()
	book := entity.Book{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
		PublishedAt: &dt,
		AuthorID:    &authorID,
	}
	if err := s.service.Create(ctx, &book); err != nil {
		if err == utils.ErrInvalidForeinKey {
			return nil, status.Errorf(codes.InvalidArgument, "author_id does not exists :%s", err)
		}
		s.logger.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to persist book:%s", err)
	}
	return converter.ConvertBook(book), nil
}

func (s *BookServer) UpdateBook(ctx context.Context, in *pb.UpdateBookRequest) (*pb.Book, error) {
	authorID := in.GetAuthorId()
	book := entity.Book{
		ID:          in.GetId(),
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
		AuthorID:    &authorID,
	}
	if err := s.service.Update(ctx, &book); err != nil {
		if err == utils.ErrNoEntity {
			return nil, status.Errorf(codes.NotFound, "failed to get book :%s", err)
		}
		s.logger.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to persist book:%s", err)
	}
	return converter.ConvertBook(book), nil
}

func (s *BookServer) DeleteBook(ctx context.Context, in *pb.PathRequest) (*pb.Empty, error) {
	if err := s.service.Delete(ctx, in.GetId()); err != nil {
		if err == utils.ErrNoEntity {
			return nil, status.Errorf(codes.NotFound, "failed to get book :%s", err)
		}
		s.logger.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to delete book:%s", err)
	}
	return nil, nil
}

func (s *BookServer) FindBook(ctx context.Context, in *pb.PathRequest) (*pb.Book, error) {
	book, err := s.service.GetByID(ctx, in.GetId())
	if err != nil {
		if err == utils.ErrNoEntity {
			return nil, status.Errorf(codes.NotFound, "failed to get book :%s", err)
		}
		s.logger.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to delete book:%s", err)
	}
	return converter.ConvertBook(book), nil
}
