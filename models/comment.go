package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
}
