package Impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/dao"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
)

type postRepositoryImpl struct {
	dao *dao.PostDAO
}

func NewPostRepository(dao *dao.PostDAO) _interface.PostRepository {
	return &postRepositoryImpl{dao: dao}
}

func (r *postRepositoryImpl) Create(ctx context.Context, post *entity.Post) error {
	return r.dao.Create(ctx, model.FromEntityPost(post))
}

func (r *postRepositoryImpl) Update(ctx context.Context, post *entity.Post) error {
	return r.dao.Update(ctx, model.FromEntityPost(post))
}

func (r *postRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.dao.Delete(ctx, id)
}

func (r *postRepositoryImpl) GetByID(ctx context.Context, id string) (*entity.Post, error) {
	modelPost, err := r.dao.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return modelPost.ToEntity(), nil
}

func (r *postRepositoryImpl) List(ctx context.Context, offset, limit int) ([]*entity.Post, error) {
	modelPosts, err := r.dao.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var entityPosts []*entity.Post
	for _, mp := range modelPosts {
		entityPosts = append(entityPosts, mp.ToEntity())
	}

	return entityPosts, nil
}

func (r *postRepositoryImpl) Search(ctx context.Context, keyword string, offset, limit int) ([]*entity.Post, error) {
	modelPosts, err := r.dao.Search(ctx, keyword, offset, limit)
	if err != nil {
		return nil, err
	}

	var entityPosts []*entity.Post
	for _, mp := range modelPosts {
		entityPosts = append(entityPosts, mp.ToEntity())
	}

	return entityPosts, nil
}
