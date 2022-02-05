package model

type UserID string
type Email string

type User struct {
	ID    UserID
	Name  string
	Email Email
}

func (u *User) SetName(name string) {
	u.Name = name
}
