package main

import (
	"database/sql"
	"github.com/tzmfreedom/go-sample/clean_architecture/controller"
	"github.com/tzmfreedom/go-sample/clean_architecture/database"
	"github.com/tzmfreedom/go-sample/clean_architecture/usecase"
	"net/http"

)

func main() {
	db, err := sql.Open("mysql", "root:password@localhost:8000")
	if err != nil {
		panic(err)
	}
	postRepo := database.NewPostRepository(db)
	useCase := usecase.NewUseCaseGetPosts(postRepo)
	controller := controller.NewController(useCase)
	http.HandleFunc("/posts", controller.GetPosts)
	http.ListenAndServe(":8080", nil)
}
