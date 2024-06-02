package repository

import "deli-blog/model"

type UserRepository interface {
	GetById(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	GetAllPosts(user *model.User) (*[]model.Post, error)
	AddPost(user *model.User, post *model.Post) error
}
