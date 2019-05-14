package models

type User struct {
	Email   string
	IsAdmin bool
}

func (u *User) GetIsAdmin() bool {
	if u == nil {
		return false
	}
	return u.IsAdmin
}

func (u *User) GetEmail() string {
	if u == nil {
		return ""
	}
	return u.Email
}
