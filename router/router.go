package router

import (
	"github.com/XMatrixStudio/IceCream/service"
	"github.com/gin-gonic/gin"
)

// Router 路由表
func Router() *gin.Engine {
	r := gin.Default()
	Comments := r.Group("/comment")
	Comments.GET("/", service.GetComment)
	return r
}
