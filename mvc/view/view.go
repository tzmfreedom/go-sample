package view

import (
	"html/template"
	"io"

	"github.com/tzmfreedom/go-sample/mvc/model"
)

type UserView struct{}

func NewUserView() *UserView {
	return &UserView{}
}

func (v *UserView) Render(w io.Writer, u *model.User) error {
	t := template.New("hoge")
	t, err := t.Parse("hello {{ .Name }}")
	if err != nil {
		return err
	}
	return t.Execute(w, u)
}
