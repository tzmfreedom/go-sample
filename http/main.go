package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/foo", &Handler{})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path)
}

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Foo %s", r.URL.Path)
}
