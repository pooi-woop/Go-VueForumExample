package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB 全局数据库实例
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("./forum.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate() error {
	return DB.AutoMigrate(&User{}, &Post{}, &Comment{})
}
