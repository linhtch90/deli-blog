package service

import (
	"deli-blog/repository"
)

type Service struct {
	user repository.UserRepository
	post repository.PostRepository
}

func NewService(user repository.UserRepository, post repository.PostRepository) *Service {
	return &Service{
		user: user,
		post: post,
	}
}
