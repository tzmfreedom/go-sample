package model

type UserID string

type User struct {
	ID   UserID
	Name string
}

func (u *User) SetName(name string) {
	u.Name = name
}
