package router

import "github.com/gin-gonic/gin"

// Router 路由表
func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"state": "success"})
	})
	return r
}
