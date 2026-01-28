package mapper

import (
	pb "github.com/KaminurOrynbek/BiznesAsh/ContentService/auto-proto/content"
	"github.com/KaminurOrynbek/BiznesAsh/ContentService/internal/entity"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertCommentToPB(c *entity.Comment) *pb.Comment {
	return &pb.Comment{
		Id:        c.ID,
		PostId:    c.PostID,
		AuthorId:  c.AuthorID,
		Content:   c.Content,
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(c.UpdatedAt),
	}
}
