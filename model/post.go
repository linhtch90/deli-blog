package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title  string `gorm:"not null"`
	Body   string
	UserID uint
}
