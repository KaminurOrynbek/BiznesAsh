package entity

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Content struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	OwnerID string `json:"owner_id"`
}

type Notification struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}
