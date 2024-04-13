package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID           uint
	ContentText      string
	ContentImagePath string
	Visible          bool
	Comments         []*Comment `gorm:"foreignKey:PostID"`
	Likes            []*User    `gorm:"many2many:post_likes;"`
}
