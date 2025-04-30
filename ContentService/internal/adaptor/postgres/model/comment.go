package model

import "time"

type Comment struct {
	ID        string    `db:"id"`
	PostID    string    `db:"post_id"`
	AuthorID  string    `db:"author_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (Comment) TableName() string {
	return "comments"
}
