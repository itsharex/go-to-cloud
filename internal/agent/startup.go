package agent

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-to-cloud/conf"
	"go-to-cloud/docs"
	"go-to-cloud/internal/middlewares"
	"go-to-cloud/internal/pkg/response"
	"io"
	"net/http"
)

func Startup() (routers *gin.Engine) {
	routers = gin.New()

	// 生产模式配置
	if conf.Environment.IsProduction() {
		gin.SetMode(gin.ReleaseMode)   // 生产模式
		gin.DefaultWriter = io.Discard // 禁用 gin 输出接口访问日志
	}

	// 开发模式配置
	if conf.Environment.IsDevelopment() {
		gin.SetMode(gin.DebugMode) // 调试模式
		buildSwagger(routers)
	}

	buildCommands(routers)

	routers.NoRoute(func(ctx *gin.Context) {
		response.GetResponse().SetHttpCode(http.StatusNotFound).FailCode(ctx, http.StatusNotFound)
	})

	return routers
}

// buildSwagger 创建swagger文档
func buildSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func buildCommands(router *gin.Engine) {

	api := router.Group("/commands")
	api.HEAD("/healthz", Healthz)

	api.Use(middlewares.AgentAuthHandler)
	{
	}
}
