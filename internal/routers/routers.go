package routers

import (
	"github.com/gin-gonic/gin"
	idx "go-to-cloud/internal/controllers/index"
	"go-to-cloud/internal/controllers/users"
)

// buildRouters 构建路由表
func buildRouters(router *gin.Engine) {
	// 登录
	index := router.Group("/index")
	{
		index.GET("/jump", idx.Jump)
	}

	user := router.Group("/user")
	{
		user.POST("/login", users.Login)
		user.GET("/info", users.Info)
		user.POST("/logout", users.Logout)
	}
}
