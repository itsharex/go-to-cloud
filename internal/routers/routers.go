package routers

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/users"
)

// buildRouters 构建路由表
func buildRouters(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.POST("/login", users.Login)
		user.GET("/info", users.Info)
		user.POST("/logout", users.Logout)
	}
}
