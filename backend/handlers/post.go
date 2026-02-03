package handlers

import (
	"net/http"
	"strconv"

	"eyu-delta-force-forum/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePostRequest 创建帖子请求结构体
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required"`
}

// UpdatePostRequest 更新帖子请求结构体
type UpdatePostRequest struct {
	Title   string `json:"title" binding:"required,min=3,max=100"`
	Content string `json:"content" binding:"required"`
}

// GetPosts 获取帖子列表
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := models.DB.Preload("User").Order("created_at DESC").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

// GetPost 获取单个帖子
func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	if err := models.DB.Preload("User").Preload("Comments").Preload("Comments.User").First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		}
		return
	}

	// 增加浏览量
	models.DB.Model(&post).Update("views", post.Views+1)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// CreatePost 创建帖子
func CreatePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID.(uint),
	}

	if err := models.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// 重新加载帖子以获取关联数据
	models.DB.Preload("User").First(&post, post.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully",
		"post":    post,
	})
}

// UpdatePost 更新帖子
func UpdatePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// 检查帖子是否存在
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		}
		return
	}

	// 检查权限
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to update this post"})
		return
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 更新帖子
	post.Title = req.Title
	post.Content = req.Content

	if err := models.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	// 重新加载帖子以获取关联数据
	models.DB.Preload("User").First(&post, post.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
		"post":    post,
	})
}

// DeletePost 删除帖子
func DeletePost(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// 检查帖子是否存在
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get post"})
		}
		return
	}

	// 检查权限
	if post.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this post"})
		return
	}

	// 删除帖子
	if err := models.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
