package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	HashedPassword string `gorm:"size:100;not null"`
	Salt           string `gorm:"size:20;not null"`
	FirstName      string `gorm:"size:50"`
	LastName       string `gorm:"size:50"`
	DOB            time.Time
	Email          string  `gorm:"size:50;unique;not null"`
	Username       string  `gorm:"size:50;not null;unique;index:idx_username"`
	Posts          []*Post `gorm:"foreignKey:UserID"`
	Followers      []*User `gorm:"many2many:user_followers;joinForeignKey:UserID;JoinReferences:FollowerID"`
	Following      []*User `gorm:"many2many:user_followings;joinForeignKey:FollowerID;JoinReferences:UserID"`
}
