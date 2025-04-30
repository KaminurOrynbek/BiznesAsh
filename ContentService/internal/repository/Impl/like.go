package Impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adaptor/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adaptor/postgres/model"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
)

type likeRepositoryImpl struct {
	dao *dao.LikeDAO
}

func NewLikeRepository(dao *dao.LikeDAO) _interface.LikeRepository {
	return &likeRepositoryImpl{dao: dao}
}

func (r *likeRepositoryImpl) Like(ctx context.Context, like *entity.Like) error {
	modelLike := model.FromEntityLike(like)
	return r.dao.Like(ctx, modelLike)
}

func (r *likeRepositoryImpl) Dislike(ctx context.Context, like *entity.Like) error {
	modelDislike := model.FromEntityLike(like)
	return r.dao.Dislike(ctx, modelDislike)
}

func (r *likeRepositoryImpl) CountLikes(ctx context.Context, postID string) (int32, error) {
	return r.dao.CountLikes(ctx, postID)
}

func (r *likeRepositoryImpl) CountDislikes(ctx context.Context, postID string) (int32, error) {
	return r.dao.CountDislikes(ctx, postID)
}
