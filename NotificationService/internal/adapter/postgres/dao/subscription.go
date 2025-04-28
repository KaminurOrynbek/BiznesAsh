package dao

type Subscription struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	EventType string `db:"event_type"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}
