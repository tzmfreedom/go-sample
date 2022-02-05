package handler

import (
	"net/http"

	"github.com/tzmfreedom/go-sample/mvc/service"

	"github.com/tzmfreedom/go-sample/mvc/model"
	"github.com/tzmfreedom/go-sample/mvc/view"

	"github.com/tzmfreedom/go-sample/mvc/repository"
)

type UserHandler struct {
	userRepo   *repository.UserRepository
	notifyRepo *repository.UserNotifyRepository
	view       *view.View
}

func NewUserHandler(userRepo *repository.UserRepository, notifyRepo *repository.UserNotifyRepository, view *view.View) *UserHandler {
	return &UserHandler{
		userRepo:   userRepo,
		notifyRepo: notifyRepo,
		view:       view,
	}
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		err := h.find(w, r)
		if err != nil {
		}
	case http.MethodPost:
		err := h.create(w, r)
		if err != nil {
		}
	case http.MethodPut:
		err := h.update(w, r)
		if err != nil {
		}
	}
}

func (h *UserHandler) find(w http.ResponseWriter, r *http.Request) error {
	vs := r.URL.Query()
	id := vs.Get("id")
	user, err := h.userRepo.FindByID(model.UserID(id))
	if err != nil {
		return err
	}
	return h.view.RenderHTML(w, "", user)
}

func (h *UserHandler) create(w http.ResponseWriter, r *http.Request) error {
	vs := r.URL.Query()
	id := vs.Get("id")
	s := service.NewUserService(h.userRepo, h.notifyRepo)
	user, err := s.UpdateAndNotify(model.UserID(id))
	if err != nil {
		return err
	}
	return h.view.RenderJSON(w, user)
}

func (h *UserHandler) update(w http.ResponseWriter, r *http.Request) error {
	vs := r.URL.Query()
	id := vs.Get("id")
	user, err := h.userRepo.FindByID(model.UserID(id))
	if err != nil {
	}
	user.SetName(vs.Get("name"))
	return h.userRepo.Update(user)
}
