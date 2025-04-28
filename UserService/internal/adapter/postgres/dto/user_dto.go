package dto

import (
	"time"

	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity/enum"
)

// UserDTO используется для маппинга данных с базой данных
type UserDTO struct {
	ID        string    `db:"id"`
	Email     string    `db:"email"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Role      string    `db:"role"` // Используем string, так как база данных хранит роль как текст
	Bio       string    `db:"bio"`
	Banned    bool      `db:"banned"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// ToUserDTO конвертирует entity.User в UserDTO
func ToUserDTO(user *entity.User) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Role:      string(user.Role), // Преобразуем enum.Role в string
		Bio:       user.Bio,
		Banned:    user.Banned,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ToEntityUser конвертирует UserDTO в entity.User
func ToEntityUser(dto *UserDTO) *entity.User {
	return &entity.User{
		ID:        dto.ID,
		Email:     dto.Email,
		Username:  dto.Username,
		Password:  dto.Password,
		Role:      enum.Role(dto.Role), // Преобразуем string в enum.Role
		Bio:       dto.Bio,
		Banned:    dto.Banned,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
