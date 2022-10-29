package routers

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/auth"
	"go-to-cloud/internal/controllers/configure/artifact"
	"go-to-cloud/internal/controllers/configure/deploy/k8s"
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
		confCodeRepo.PUT("/", scm.UpdateCodeRepo)
		confCodeRepo.DELETE("/:id", scm.RemoveCodeRepo)
		confCodeRepo.POST("/testing", scm.Testing)

		confArtifact := api.Group("/configure/artifact")
		confArtifact.POST("/testing", artifact.Testing)
		confArtifact.POST("/bind", artifact.BindArtifactRepo)
		confArtifact.PUT("/", artifact.UpdateArtifactRepo)
		confArtifact.GET("/", artifact.QueryArtifactRepo)
		confArtifact.DELETE("/:id", artifact.RemoveArtifactRepo)
		confArtifact.GET("/:id", artifact.QueryArtifactItems)

		confK8s := api.Group("/configure/deploy/k8s")
		confK8s.POST("/testing", k8s.Testing)
		confK8s.POST("/bind", k8s.BindK8sRepo)
		confK8s.PUT("/", k8s.UpdateK8sRepo)
		confK8s.GET("/", k8s.QueryK8sRepos)
		confK8s.DELETE("/:id", k8s.RemoveK8sRepo)
		confK8s.GET("/:id", k8s.QueryK8sRepos)

		project := api.Group("/projects")
		project.GET("/list", projects.List)
		project.GET("/coderepo", projects.CodeRepo)
	}
}
