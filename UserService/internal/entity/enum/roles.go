package enum

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleModerator Role = "moderator"
	RoleUser      Role = "user"
)

func (r Role) IsAdmin() bool {
	return r == RoleAdmin
}

func (r Role) IsModerator() bool {
	return r == RoleModerator
}

func (r Role) IsUser() bool {
	return r == RoleUser
}
