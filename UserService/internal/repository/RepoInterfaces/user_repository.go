package RepoInterfaces

import (
	"context"
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/entity"
)

// UserFilter contains filtering options for listing users
type UserFilter struct {
	Email    string
	Username string
	Role     string
	Banned   *bool
	Limit    int
	Offset   int
}

// UserRepository defines the interface for user data operations
type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, filter entity.UserFilter) ([]*entity.User, error)
	BanUser(ctx context.Context, id string) error
}
