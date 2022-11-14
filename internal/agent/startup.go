package agent

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/conf"
	"go-to-cloud/internal/controllers/users"
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
	}

	buildCommands(routers)

	routers.NoRoute(func(ctx *gin.Context) {
		response.GetResponse().SetHttpCode(http.StatusNotFound).FailCode(ctx, http.StatusNotFound)
	})

	return routers
}

func buildCommands(router *gin.Engine) {

	api := router.Group("/commands")

	api.Use(middlewares.AgentAuthHandler)
	{
		api.POST("/logout", users.Logout)
	}
}
