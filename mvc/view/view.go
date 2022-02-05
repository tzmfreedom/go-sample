package view

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
)

type View struct{}

func NewView() *View {
	return &View{}
}

func (v *View) RenderHTML(w io.Writer, f string, data interface{}) error {
	t, err := template.ParseFiles(f)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}

func (v *View) RenderJSON(w io.Writer, data interface{}) error {
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, string(buf))
	return err
}
