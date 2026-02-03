package handlers

import (
	"net/http"

	"eyu-delta-force-forum/models"

	"github.com/gin-gonic/gin"
)

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
