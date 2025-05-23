package _interface

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
)

type LikeRepository interface {
	Like(ctx context.Context, like *entity.Like) (int32, error)
	Dislike(ctx context.Context, like *entity.Like) (int32, error)
	CountLikes(ctx context.Context, postID string) (int32, error)
	CountDislikes(ctx context.Context, postID string) (int32, error)
	LikeComment(ctx context.Context, like *entity.Like) (int32, error)
	CountCommentLikes(ctx context.Context, commentID string) (int32, error)
}
