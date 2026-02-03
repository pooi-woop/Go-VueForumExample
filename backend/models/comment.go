package models

import (
	"time"

	"gorm.io/gorm"
)

// Comment 评论模型
type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	UserID    uint           `json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	PostID    uint           `json:"post_id"`
	Post      Post           `gorm:"foreignKey:PostID" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
