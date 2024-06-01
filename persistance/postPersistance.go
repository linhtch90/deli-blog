package persistance

import (
	"deli-blog/model"

	"gorm.io/gorm"
)

type PostPersistance struct {
	db *gorm.DB
}

func NewPostPersistance(db *gorm.DB) *PostPersistance {
	return &PostPersistance{
		db: db,
	}
}

func (postPersistance *PostPersistance) GetById(id uint) (*model.Post, error) {
	var post model.Post
	if err := postPersistance.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (postPersistance *PostPersistance) GetByTitle(title string) (*model.Post, error) {
	var post model.Post
	if err := postPersistance.db.Where(&model.Post{Title: title}).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (postPersistance *PostPersistance) Create(post *model.Post) error {
	return postPersistance.db.Create(post).Error
}

func (postPersistance *PostPersistance) Update(post *model.Post) error {
	return postPersistance.db.Model(post).Updates(post).Error
}
