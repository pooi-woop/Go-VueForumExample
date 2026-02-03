package models

import (
	"time"

	"gorm.io/gorm"
)

// Post 帖子模型
type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:100;not null" json:"title"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	UserID    uint           `json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Views     int            `gorm:"default:0" json:"views"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Comments  []Comment      `gorm:"foreignKey:PostID" json:"comments"`
}
