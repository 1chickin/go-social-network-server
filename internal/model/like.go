package model

import (
	"time"
)

type Like struct {
	PostID    uint `gorm:"primaryKey;autoIncrement:false"`
	UserID    uint `gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time
}
