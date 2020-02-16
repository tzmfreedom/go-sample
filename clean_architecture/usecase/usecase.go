package usecase

import (
	"github.com/tzmfreedom/go-sample/clean_architecture/database"
	"github.com/tzmfreedom/go-sample/clean_architecture/domain"
)

type UseCaseGetPosts struct {
	postRepo PostRepository
}

func NewUseCaseGetPosts(postRepo PostRepository) *UseCaseGetPosts {
	return &UseCaseGetPosts{postRepo}
}

func (u *UseCaseGetPosts) Run(name string) []*domain.Post {
	items := u.postRepo.GetPosts(name)
	posts := make([]*domain.Post, len(items))
	for i, item := range items {
		posts[i] = &domain.Post{
			Title: item.Title,
			Body:  item.Body,
		}
	}
	return posts
}

type PostRepository interface {
	GetPosts(string) []*database.PostDto
}
