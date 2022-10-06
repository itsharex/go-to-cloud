package routers

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/auth"
	"go-to-cloud/internal/controllers/configure/scm"
	"go-to-cloud/internal/controllers/projects"
	"go-to-cloud/internal/controllers/users"
	"go-to-cloud/internal/middlewares"
)

// buildRouters 构建路由表
func buildRouters(router *gin.Engine) {

	api := router.Group("/api")

	api.POST("/login", auth.Login)

	api.Use(middlewares.AuthHandler())
	{
		user := api.Group("/user")
		user.GET("/info", users.Info)
		user.POST("/logout", users.Logout)

		conf := api.Group("/configure")
		conf.GET("/coderepo", scm.QueryCodeRepos)
		conf.POST("/coderepo/bind", scm.BindCodeRepo)
		conf.PUT("/coderepo/bind", scm.UpdateCodeRepo)
		conf.DELETE("/coderepo/:id", scm.RemoveCodeRepo)
		conf.POST("/coderepo/testing", scm.Testing)

		project := api.Group("/projects")
		project.GET("/list", projects.List)
	}
}
