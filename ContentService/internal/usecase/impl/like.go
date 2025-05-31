package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/nats/payloads"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/nats/publisher"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	usecase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"github.com/google/uuid"
	"time"
)

type likeUsecaseImpl struct {
	likeRepo         _interface.LikeRepository
	contentPublisher *publisher.ContentPublisher
}

func NewLikeUsecase(likeRepo _interface.LikeRepository, contentPublisher *publisher.ContentPublisher) usecase.LikeUsecase {
	return &likeUsecaseImpl{
		likeRepo:         likeRepo,
		contentPublisher: contentPublisher,
	}
}

func (u *likeUsecaseImpl) LikePost(ctx context.Context, like *entity.Like) (int32, error) {
	like.ID = uuid.NewString()
	like.IsLike = true
	like.CreatedAt = time.Now()

	count, err := u.likeRepo.Like(ctx, like)
	if err != nil {
		return 0, err
	}
	_ = u.contentPublisher.PublishPostLiked(payloads.PostLiked{
		UserID: like.UserID,
		PostID: like.PostID,
	})

	return count, nil

}

func (u *likeUsecaseImpl) DislikePost(ctx context.Context, like *entity.Like) (int32, error) {
	like.ID = uuid.NewString()
	like.IsLike = false
	like.CreatedAt = time.Now()
	return u.likeRepo.Dislike(ctx, like)
}

func (u *likeUsecaseImpl) LikeComment(ctx context.Context, like *entity.Like) (int32, error) {
	like.ID = uuid.NewString()
	like.IsLike = true
	like.CreatedAt = time.Now()

	count, err := u.likeRepo.LikeComment(ctx, like)
	if err != nil {
		return 0, err
	}

	//Publish event to NotificationService
	_ = u.contentPublisher.PublishCommentLiked(payloads.CommentLiked{
		UserID:    like.UserID,
		CommentID: like.CommentID,
	})

	return count, nil
}
