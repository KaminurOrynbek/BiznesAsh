package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	ucase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"

	"time"
)

type postUsecaseImpl struct {
	postRepo _interface.PostRepository
}

func NewPostUsecase(postRepo _interface.PostRepository) ucase.PostUsecase {
	return &postUsecaseImpl{postRepo: postRepo}
}

func (u *postUsecaseImpl) CreatePost(ctx context.Context, post *entity.Post) error {
	post.CreatedAt = time.Now()
	post.UpdatedAt = post.CreatedAt
	return u.postRepo.Create(ctx, post)
}

func (u *postUsecaseImpl) UpdatePost(ctx context.Context, post *entity.Post) error {
	post.UpdatedAt = time.Now()
	return u.postRepo.Update(ctx, post)
}

func (u *postUsecaseImpl) DeletePost(ctx context.Context, id string) error {
	return u.postRepo.Delete(ctx, id)
}

func (u *postUsecaseImpl) GetPost(ctx context.Context, id string) (*entity.Post, error) {
	return u.postRepo.GetByID(ctx, id)
}

func (u *postUsecaseImpl) ListPosts(ctx context.Context) ([]*entity.Post, error) {
	// You could later improve it by adding pagination (offset, limit)
	return u.postRepo.List(ctx, 0, 100)
}

func (u *postUsecaseImpl) SearchPosts(ctx context.Context, keyword string) ([]*entity.Post, error) {
	return u.postRepo.Search(ctx, keyword, 0, 100)
}
