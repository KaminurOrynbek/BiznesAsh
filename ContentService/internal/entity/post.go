package entity

import "time"

type PostType string

const (
	PostTypeLegalInfo PostType = "LEGAL_INFO"
	PostTypeGuide     PostType = "GUIDE"
)

type Post struct {
	ID            string
	Title         string
	Content       string
	Type          PostType
	AuthorID      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Published     bool
	LikesCount    int32
	DislikesCount int32
}
