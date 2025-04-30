package entity

type UserFilter struct {
	Email       string
	Username    string
	Role        string
	Banned      *bool
	Limit       int
	Offset      int
	SearchQuery string
}
