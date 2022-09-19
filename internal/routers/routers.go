package routers

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/auth"
	"go-to-cloud/internal/controllers/users"
	"go-to-cloud/internal/middlewares"
)

// buildRouters 构建路由表
func buildRouters(router *gin.Engine) {

	router.POST("/login", auth.Login)

	api := router.Group("/api")
	api.Use(middlewares.AuthHandler())
	{
		user := api.Group("/user")
		user.GET("/info", users.Info)
		user.POST("/logout", users.Logout)
	}
}
