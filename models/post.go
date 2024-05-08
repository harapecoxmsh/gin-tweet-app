package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content string `gorm:"not null"`
}
