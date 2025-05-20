package impl

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	_interface "github.com/KaminurOrynbek/BiznesAsh/internal/repository/interface"
	usecase "github.com/KaminurOrynbek/BiznesAsh/internal/usecase/interface"
	"github.com/google/uuid"

	"time"
)

type postUsecaseImpl struct {
	postRepo    _interface.PostRepository
	commentRepo _interface.CommentRepository
	likeRepo    _interface.LikeRepository
}

func NewPostUsecase(postRepo _interface.PostRepository, commentRepo _interface.CommentRepository, likeRepo _interface.LikeRepository) usecase.PostUsecase {
	return usecase.PostUsecase(&postUsecaseImpl{
		postRepo:    postRepo,
		commentRepo: commentRepo,
		likeRepo:    likeRepo,
	})
}

func (u *postUsecaseImpl) CreatePost(ctx context.Context, post *entity.Post) error {
	post.ID = uuid.NewString()
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
	post, err := u.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	comments, err := u.commentRepo.ListByPostID(ctx, id)
	if err != nil {
		return nil, err
	}
	var valComments []*entity.Comment
	for _, c := range comments {
		valComments = append(valComments, c)
	}
	post.Comments = valComments

	post.CommentsCount = int32(len(comments))

	likes, _ := u.likeRepo.CountLikes(ctx, id)
	dislikes, _ := u.likeRepo.CountDislikes(ctx, id)
	post.LikesCount = likes
	post.DislikesCount = dislikes

	return post, nil
}

func (u *postUsecaseImpl) ListPosts(ctx context.Context, offset, limit int) ([]*entity.Post, error) {
	posts, err := u.postRepo.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		comments, err := u.commentRepo.ListByPostID(ctx, post.ID)
		if err == nil {
			post.Comments = comments
			post.CommentsCount = int32(len(comments))
		}

		likes, _ := u.likeRepo.CountLikes(ctx, post.ID)
		dislikes, _ := u.likeRepo.CountDislikes(ctx, post.ID)
		post.LikesCount = likes
		post.DislikesCount = dislikes
	}

	return posts, nil
}

func (u *postUsecaseImpl) SearchPosts(ctx context.Context, keyword string, offset, limit int) ([]*entity.Post, error) {
	posts, err := u.postRepo.Search(ctx, keyword, offset, limit)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		comments, err := u.commentRepo.ListByPostID(ctx, post.ID)
		if err != nil {
			continue
		}
		post.CommentsCount = int32(len(comments))
	}

	return posts, nil
}
