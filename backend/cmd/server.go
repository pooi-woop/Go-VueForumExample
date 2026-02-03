package main

import (
	"log"
	"os"

	"eyu-delta-force-forum/handlers"
	"eyu-delta-force-forum/middleware"
	"eyu-delta-force-forum/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 初始化数据库
	if err := models.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 自动迁移数据库
	if err := models.AutoMigrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode)

	// 创建Gin引擎
	r := gin.Default()

	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API路由组
	api := r.Group("/api")
	{
		// 认证路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		// 帖子路由
		posts := api.Group("/posts")
		{
			posts.GET("", handlers.GetPosts)
			posts.GET("/:id", handlers.GetPost)
			posts.POST("", middleware.JWTAuth(), handlers.CreatePost)
			posts.PUT("/:id", middleware.JWTAuth(), handlers.UpdatePost)
			posts.DELETE("/:id", middleware.JWTAuth(), handlers.DeletePost)
		}

		// 评论路由
		comments := api.Group("/comments")
		{
			comments.GET("/post/:postId", handlers.GetComments)
			comments.POST("", middleware.JWTAuth(), handlers.CreateComment)
			comments.DELETE("/:id", middleware.JWTAuth(), handlers.DeleteComment)
		}

		// 用户路由
		users := api.Group("/users")
		{
			users.GET("/me", middleware.JWTAuth(), handlers.GetCurrentUser)
		}
	}

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
