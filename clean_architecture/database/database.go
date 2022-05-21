package database

import (
	"database/sql"
)

type PostDto struct {
	Title string
	Body  string
}

type PostRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{db}
}

func (r *PostRepositoryImpl) GetPosts(name string) []*PostDto {
	return []*PostDto{
		{
			Title: "Hoge",
			Body:  "Fuga",
		},
	}
}
