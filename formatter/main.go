package main

import "fmt"

type Hoge struct {
	Name string
	Id   int
}

func (h *Hoge) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		fmt.Fprintf(s, "Hoge => Id: %d, Name: %s", h.Id, h.Name)
	case 'n':
		fmt.Fprintf(s, "Hoge.Name => %s", h.Name)
	case 'd':
		fmt.Fprintf(s, "Hoge.Id => %d", h.Id)
	}
}

func main() {
	fmt.Printf("%s\n", &Hoge{Name: "hoge", Id: 123})
	fmt.Printf("%d\n", &Hoge{Name: "hoge", Id: 123})
	fmt.Printf("%n\n", &Hoge{Name: "hoge", Id: 123})
}
