package converter

import (
	"example/grpc/internal/controller/rpc/pb"
	"example/grpc/internal/core/entity"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertAuthor(author entity.Author) *pb.Author {
	b := &pb.Author{
		Id:        author.ID,
		Name:      author.Name,
		Biography: author.Biography,
	}
	if author.CreatedAt != nil {
		b.CreatedAt = timestamppb.New(*author.CreatedAt)
	}
	if author.UpdatedAt != nil {
		b.UpdatedAt = timestamppb.New(*author.UpdatedAt)
	}
	return b
}

func ConvertListAuthor(authors []entity.Author) []*pb.Author {
	result := make([]*pb.Author, 0, cap(authors))
	for _, author := range authors {
		result = append(result, ConvertAuthor(author))
	}
	return result
}
