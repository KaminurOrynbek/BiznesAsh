package dao

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/internal/adapter/postgres/model"
	"github.com/jmoiron/sqlx"
)

type PostDAO struct {
	db *sqlx.DB
}

func NewPostDAO(db *sqlx.DB) *PostDAO {
	return &PostDAO{db: db}
}

func (dao *PostDAO) Create(ctx context.Context, post *model.Post) error {
	query := `
		INSERT INTO posts (id, title, content, type, author_id, created_at, updated_at, published, likes_count, dislikes_count, comments_count)
		VALUES (:id, :title, :content, :type, :author_id, :created_at, :updated_at, :published, :likes_count, :dislikes_count, :comments_count)
	`
	_, err := dao.db.NamedExecContext(ctx, query, post)
	return err
}

func (dao *PostDAO) Update(ctx context.Context, post *model.Post) error {
	query := `
		UPDATE posts SET 
			title = :title, 
			content = :content, 
			type = :type, 
			updated_at = :updated_at, 
			published = :published,
			comments_count = :comments_count
		WHERE id = :id
	`
	_, err := dao.db.NamedExecContext(ctx, query, post)
	return err
}

func (dao *PostDAO) Delete(ctx context.Context, postID string) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := dao.db.ExecContext(ctx, query, postID)
	return err
}

func (dao *PostDAO) GetByID(ctx context.Context, id string) (*model.Post, error) {
	query := `
		SELECT id, title, content, type, author_id, created_at, updated_at, published, likes_count, dislikes_count, comments_count
		FROM posts
		WHERE id = $1
	`
	var post model.Post
	if err := dao.db.GetContext(ctx, &post, query, id); err != nil {
		return nil, err
	}

	// Manually preload comments
	commentQuery := `
		SELECT id, post_id, author_id, content, created_at, updated_at
		FROM comments
		WHERE post_id = $1
	`
	var comments []*model.Comment
	if err := dao.db.SelectContext(ctx, &comments, commentQuery, id); err != nil {
		return nil, err
	}
	post.Comments = comments

	return &post, nil
}

func (dao *PostDAO) List(ctx context.Context, offset, limit int) ([]*model.Post, error) {
	query := `
		SELECT id, title, content, type, author_id, created_at, updated_at, published, likes_count, dislikes_count, comments_count
		FROM posts
		ORDER BY created_at DESC
	`
	var posts []*model.Post
	err := dao.db.SelectContext(ctx, &posts, query)
	return posts, err
}

func (dao *PostDAO) Search(ctx context.Context, keyword string, offset, limit int) ([]*model.Post, error) {
	query := `
		SELECT id, title, content, type, author_id, created_at, updated_at, published, likes_count, dislikes_count, comments_count
		FROM posts
		WHERE title ILIKE '%' || $1 || '%' OR content ILIKE '%' || $1 || '%'
		ORDER BY created_at DESC
		OFFSET $2 LIMIT $3
	`
	var posts []*model.Post
	err := dao.db.SelectContext(ctx, &posts, query, keyword, offset, limit)
	return posts, err
}
