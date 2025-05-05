package Impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
)

type commentRepositoryImpl struct {
	dao *dao.CommentDAO
}

func NewCommentRepository(dao *dao.CommentDAO) _interface.CommentRepository {
	return &commentRepositoryImpl{dao: dao}
}

func (r *commentRepositoryImpl) Create(ctx context.Context, comment *entity.Comment) error {
	modelComment := model.FromEntityComment(comment)
	return r.dao.Create(ctx, modelComment)
}

func (r *commentRepositoryImpl) Update(ctx context.Context, comment *entity.Comment) error {
	modelComment := model.FromEntityComment(comment)
	return r.dao.Update(ctx, modelComment)
}

func (r *commentRepositoryImpl) Delete(ctx context.Context, commentID string) error {
	return r.dao.Delete(ctx, commentID)
}

func (r *commentRepositoryImpl) ListByPostID(ctx context.Context, postID string) ([]*entity.Comment, error) {
	modelComments, err := r.dao.ListByPostID(ctx, postID)
	if err != nil {
		return nil, err
	}

	var entityComments []*entity.Comment
	for _, m := range modelComments {
		entityComments = append(entityComments, m.ToEntity())
	}
	return entityComments, nil
}
