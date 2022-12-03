package controller

import (
	"context"
	pb "example/grpc/pb/service"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var authors = make([]*pb.Author, 0)

// AuthorServer serves gRPC requests for core business logics services.
type AuthorController struct {
	pb.UnimplementedAuthorServiceServer
}

// NewServer created a new gRPC server.
func NewAuthorServer() *AuthorController {
	return &AuthorController{}
}

func findAuthor(id int64) (int, error) {
	var author = -1
	for k, v := range authors {
		if v.Id == id {
			author = k
		}
	}
	if author == -1 {
		return -1, status.Errorf(codes.NotFound, "author not found: ")
	}
	return author, nil
}

func (c *AuthorController) GetAuthors(context.Context, *pb.QueryRequest) (*pb.AuthorListResponse, error) {
	for i := 0; i < 3; i++ {
		authors = append(authors, &pb.Author{
			Name:      fmt.Sprintf("author %d", i+1),
			Id:        int64(i) + 1,
			Biography: "In publishing and graphic design, Lorem ipsum is a placeholder",
		})
	}
	return &pb.AuthorListResponse{Authors: authors}, nil
}

func (c *AuthorController) FindAuthor(ctx context.Context, req *pb.PathRequest) (*pb.Author, error) {
	index, err := findAuthor(req.Id)
	if err != nil {
		return &pb.Author{}, err
	}
	return authors[index], nil
}

func (c *AuthorController) CreateAuthor(ctx context.Context, req *pb.CreateAuthorRequest) (*pb.Author, error) {
	author := &pb.Author{
		Id:        int64(len(authors)) + 1,
		Name:      req.GetName(),
		Biography: req.GetBiography(),
	}
	authors = append(authors, author)

	return author, nil
}

func (c *AuthorController) UpdateAuthor(ctx context.Context, req *pb.UpdateAuthorRequest) (*pb.Author, error) {
	index, err := findAuthor(req.Id)
	if err != nil {
		return &pb.Author{}, err
	}
	authors[index].Name = req.GetName()
	authors[index].Biography = req.GetBiography()
	return authors[index], nil
}

func (c *AuthorController) DeleteAuthor(ctx context.Context, req *pb.PathRequest) (*pb.Empty, error) {
	index, err := findAuthor(req.Id)
	if err != nil {
		return nil, err
	}
	authors = append(authors[:index], authors[index+1:]...)
	return &pb.Empty{}, nil
}
