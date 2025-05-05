package model

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"time"
)

type Verification struct {
	ID        string    `db:"id"` // Keep as is; omit 'omitempty'
	UserID    string    `db:"user_id"`
	Email     string    `db:"email"`
	Code      string    `db:"code"`
	ExpiresAt time.Time `db:"expires_at"`
	IsUsed    bool      `db:"is_used"`
}

func (Verification) TableName() string {
	return "verifications"
}

func (v *Verification) ToEntity() *entity.Verification {
	return &entity.Verification{
		ID:        v.ID,
		UserID:    v.UserID,
		Email:     v.Email,
		Code:      v.Code,
		ExpiresAt: v.ExpiresAt,
		IsUsed:    v.IsUsed,
	}
}

func FromEntityVerification(e *entity.Verification) *Verification {
	return &Verification{
		ID:        e.ID,
		UserID:    e.UserID,
		Email:     e.Email,
		Code:      e.Code,
		ExpiresAt: e.ExpiresAt,
		IsUsed:    e.IsUsed,
	}
}
