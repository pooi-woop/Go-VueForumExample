package config

import (
	"os"
)

// Config 应用配置
type Config struct {
	JWTSecret     string
	DatabaseURL   string
	ServerPort    string
	TokenExpiry   int
}

// AppConfig 全局配置实例
var AppConfig Config

// LoadConfig 加载配置
func LoadConfig() {
	AppConfig = Config{
		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
		DatabaseURL:   getEnv("DATABASE_URL", "./forum.db"),
		ServerPort:    getEnv("PORT", "8080"),
		TokenExpiry:   24, // 24小时
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
