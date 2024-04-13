package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID  uint
	UserID  uint
	Content string
}
