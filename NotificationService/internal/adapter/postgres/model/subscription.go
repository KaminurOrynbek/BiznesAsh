package model

import (
	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"strings"
)

type Subscription struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	EventType string `db:"event_type"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}

func (s *Subscription) ToEntity() *entity.Subscription {
	return &entity.Subscription{
		ID:        s.ID,
		UserID:    s.UserID,
		EventType: strings.Split(s.EventType, ","),
	}
}

func FromEntitySubscription(e *entity.Subscription) *Subscription {
	return &Subscription{
		ID:        e.ID,
		UserID:    e.UserID,
		EventType: strings.Join(e.EventType, ","),
	}
}
