package models

import (
	"eyu-delta-force-forum/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局数据库实例
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	var err error
	config.LoadConfig()
	DB, err = gorm.Open(mysql.Open(config.AppConfig.DatabaseURL), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate() error {
	return DB.AutoMigrate(&User{}, &Post{}, &Comment{})
}
