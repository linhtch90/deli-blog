package persistance

import (
	"deli-blog/model"

	"gorm.io/gorm"
)

const ASSOCIATION_POSTS = "Posts"

type UserPersistance struct {
	db *gorm.DB
}

func NewUserPersistance(db *gorm.DB) *UserPersistance {
	return &UserPersistance{
		db: db,
	}
}

func (userPersistance *UserPersistance) GetById(id uint) (*model.User, error) {
	var user model.User
	if err := userPersistance.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userPersistance *UserPersistance) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := userPersistance.db.Where(&model.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userPersistence *UserPersistance) Create(user *model.User) error {
	return userPersistence.db.Create(user).Error
}

func (userPersistance *UserPersistance) Update(user *model.User) error {
	return userPersistance.db.Model(user).Updates(user).Error
}

func (userPersistance *UserPersistance) GetAllPosts(user *model.User) (*[]model.Post, error) {
	var posts []model.Post
	if error := userPersistance.db.Model(user).Association(ASSOCIATION_POSTS).Find(&posts); error != nil {
		return nil, error
	}
	return &posts, nil
}

func (userPersistance *UserPersistance) AddPost(user *model.User, post *model.Post) error {
	return userPersistance.db.Model(user).Association(ASSOCIATION_POSTS).Append(post)
}
