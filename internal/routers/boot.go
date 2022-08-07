package routers

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

// SetRouters 设置API路由
func SetRouters() (routers *gin.Engine) {
	routers = gin.New()

	// 生产模式配置
	if conf.Enviroment.IsProduction() {
		gin.SetMode(gin.ReleaseMode)   // 生产模式
		gin.DefaultWriter = io.Discard // 禁用 gin 输出接口访问日志

		routers.Use(
			middlewares.GenericRecovery(),
			middlewares.CorsHandler(),
		)
	}

	// 开发模式配置
	if conf.Enviroment.IsDevelopment() {
		gin.SetMode(gin.DebugMode) // 调试模式
		routers.Use(
			middlewares.GenericRecovery(),
			gin.Logger(),
			middlewares.CorsHandler(),
		)
	}

	// 构建swagger
	buildSwagger(routers)

	// 构建路由
	buildRouters(routers)

	routers.NoRoute(func(ctx *gin.Context) {
		response.GetResponse().SetHttpCode(http.StatusNotFound).FailCode(ctx, http.StatusNotFound)
	})

	return
}

// buildSwagger 创建swagger文档
func buildSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/user"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
