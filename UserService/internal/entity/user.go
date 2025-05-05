package entity

import (
	"github.com/KaminurOrynbek/BiznesAsh/UserService/internal/entity/enum"
	"time"
)

type User struct {
	ID        string
	Email     string
	Username  string
	Password  string
	Role      enum.Role
	Bio       string
	Banned    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
