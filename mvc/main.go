package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/tzmfreedom/go-sample/mvc/view"

	"github.com/tzmfreedom/go-sample/mvc/handler"
	"github.com/tzmfreedom/go-sample/mvc/repository"
)

func main() {
	db, err := sql.Open("mysql", "localhost")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}()
	c := container{db: db}
	http.Handle("/user", c.GetUserHandler())
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type container struct {
	db *sql.DB
}

func (c *container) GetUserHandler() *handler.UserHandler {
	return handler.NewUserHandler(c.GetUserRepository(), c.GetUserView())
}

func (c *container) GetUserRepository() *repository.UserRepository {
	return repository.NewUserRepository(c.db)
}

func (c *container) GetUserView() *view.UserView {
	return view.NewUserView()
}
