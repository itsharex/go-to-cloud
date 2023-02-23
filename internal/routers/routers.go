package routers

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/auth"
	"go-to-cloud/internal/controllers/configure/artifact"
	"go-to-cloud/internal/controllers/configure/buildEnv"
	"go-to-cloud/internal/controllers/configure/builder"
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

		api.GET("/configure/build/env", buildEnv.BuildEnv)
		api.GET("/configure/build/cmd", buildEnv.BuildCmd)

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

		confBuilder := api.Group("/configure/builder")
		confBuilder.POST("/install/k8s", builder.K8sInstall)
		confBuilder.GET("/nodes/k8s", builder.QueryNodesOnK8s)
		confBuilder.GET("/nodes/k8s/available", builder.QueryAvailableNodesOnK8s)
		confBuilder.DELETE("/node/:id", builder.Uninstall)
		confBuilder.PUT("/node", builder.UpdateBuilderNode)

		project := api.Group("/projects")
		project.POST("/", projects.Create)
		project.DELETE("/:projectId", projects.DeleteProject)
		project.GET("/list", projects.List)
		project.GET("/coderepo", projects.CodeRepo)
		project.PUT("/", projects.UpdateProject)
		project.POST("/:projectId/import", projects.ImportSourceCode)
		project.GET("/:projectId/imported", projects.ListImportedSourceCode)
		project.DELETE("/:projectId/sourcecode/:id", projects.DeleteSourceCode)
		project.GET("/:projectId/src/:sourceCodeId", projects.ListBranches)
		project.POST("/:projectId/pipeline", projects.NewBuildPlan)
		project.GET("/:projectId/pipeline", projects.QueryBuildPlan)
		project.GET("/:projectId/pipeline/state", projects.QueryBuildPlanState)
		project.DELETE("/:projectId/pipeline/:id", projects.DeleteBuildPlan)
		project.POST("/:projectId/pipeline/:id/build", projects.StartBuildPlan)
		project.GET("/:projectId/deploy/apps", projects.QueryDeployedApps)
		project.POST("/:projectId/deploy/app", projects.CreateDeployment)
	}
}
