package repository

import (
	"database/sql"

	"github.com/tzmfreedom/go-sample/mvc/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByID(id model.UserID) (*model.User, error) {
	row := r.db.QueryRow("SELECT name FROM users WHERE id = ?", id)
	var name string
	err := row.Scan(&name)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   id,
		Name: name,
	}, nil
}

func (r *UserRepository) Update(user *model.User) error {
	_, err := r.db.Exec("UPDATE users")
	return err
}
