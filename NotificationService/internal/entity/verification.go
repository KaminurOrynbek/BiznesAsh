package entity

import "time"

type Verification struct {
	ID        string
	UserID    string
	Email     string
	Code      string //verification code
	ExpiresAt time.Time
	IsUsed    bool //// whether the code was already used
}
