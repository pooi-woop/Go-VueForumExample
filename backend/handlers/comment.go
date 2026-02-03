package handlers

import (
	"net/http"
	"strconv"

	"eyu-delta-force-forum/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateCommentRequest 创建评论请求结构体
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}

// GetComments 获取帖子的评论列表
func GetComments(c *gin.Context) {
	postIDStr := c.Param("postId")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var comments []models.Comment
	if err := models.DB.Where("post_id = ?", postID).Preload("User").Order("created_at DESC").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查帖子是否存在
	var post models.Post
	if err := models.DB.First(&post, req.PostID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check post"})
		}
		return
	}

	comment := models.Comment{
		Content: req.Content,
		UserID:  userID.(uint),
		PostID:  req.PostID,
	}

	if err := models.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	// 重新加载评论以获取关联数据
	models.DB.Preload("User").First(&comment, comment.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment created successfully",
		"comment": comment,
	})
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	// 检查评论是否存在
	var comment models.Comment
	if err := models.DB.First(&comment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comment"})
		}
		return
	}

	// 检查权限
	if comment.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete this comment"})
		return
	}

	// 删除评论
	if err := models.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully",
	})
}
