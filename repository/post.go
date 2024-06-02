package repository

import "deli-blog/model"

type PostRepository interface {
	GetById(id uint) (*model.Post, error)
	GetByTitle(title string) (*model.Post, error)
}
