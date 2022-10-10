package routers

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/auth"
	"go-to-cloud/internal/controllers/configure/artifact"
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

		confCodeRepo := api.Group("/configure/coderepo")
		confCodeRepo.GET("/", scm.QueryCodeRepos)
		confCodeRepo.POST("/bind", scm.BindCodeRepo)
		confCodeRepo.PUT("/bind", scm.UpdateCodeRepo)
		confCodeRepo.DELETE("/:id", scm.RemoveCodeRepo)
		confCodeRepo.POST("/testing", scm.Testing)

		confArtifact := api.Group("/configure/artifact")
		confArtifact.POST("/testing", artifact.Testing)

		project := api.Group("/projects")
		project.GET("/list", projects.List)
	}
}
