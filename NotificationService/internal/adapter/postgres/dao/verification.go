package dao

import "time"

type Verification struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Email     string    `db:"email"`
	Code      string    `db:"code"`
	ExpiresAt time.Time `db:"expires_at"`
	IsUsed    bool      `db:"is_used"`
}

func (Verification) TableName() string {
	return "verifications"
}
