package persistance

import (
	"deli-blog/model"

	"gorm.io/gorm"
)

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
