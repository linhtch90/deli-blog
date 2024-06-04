package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	Address  *string
	Posts    []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
