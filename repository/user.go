package repository

import "deli-blog/model"

type UserRepository interface {
	GetById(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
}
