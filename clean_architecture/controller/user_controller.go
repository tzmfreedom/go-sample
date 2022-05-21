package controller

import (
	"net/http"

	"github.com/tzmfreedom/go-sample/clean_architecture/usecase"
)

type Controller struct {
	useCase *usecase.UseCaseGetPosts
}

func NewController(useCase *usecase.UseCaseGetPosts) *Controller {
	return &Controller{useCase}
}

func (c *Controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")
	c.useCase.Run(name)
}
