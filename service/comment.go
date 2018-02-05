package service

import "github.com/gin-gonic/gin"

// GetComment 获取评论
func GetComment(c *gin.Context) {
	c.JSON(200, gin.H{"state": "success!"})
}
