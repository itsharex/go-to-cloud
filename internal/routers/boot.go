package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-to-cloud/conf"
	"go-to-cloud/docs"
	"go-to-cloud/internal/agent"
	"go-to-cloud/internal/middlewares"
	"go-to-cloud/internal/pkg/response"
	"io"
	"net/http"
	"net/url"
)

func StartAgentServer() {
	addr := conf.GetServerGrpcHost()
	if u, err := url.Parse(addr.Host); err != nil {
		panic(err)
	} else {
		port := u.Port()
		if len(port) == 0 {
			panic(fmt.Sprintf("grpc port not found in server.host：%s", addr.Host))
		}

		err = agent.Startup(&port)
		if err != nil {
			panic(err)
		}
	}
}

// SetRouters 设置API路由
func SetRouters() (routers *gin.Engine) {
	routers = gin.New()

	// 中间件
	handlers := make([]gin.HandlerFunc, 0)
	handlers = append(handlers, middlewares.GenericRecovery())
	if conf.Environment.IsDevelopment() {
		handlers = append(handlers, gin.Logger())
	}
	handlers = append(handlers, middlewares.CorsHandler())

	// 生产模式配置
	if conf.Environment.IsProduction() {
		gin.SetMode(gin.ReleaseMode)   // 生产模式
		gin.DefaultWriter = io.Discard // 禁用 gin 输出接口访问日志
	}

	// 开发模式配置
	if conf.Environment.IsDevelopment() {
		gin.SetMode(gin.DebugMode) // 调试模式
		buildSwagger(routers)      // 构建swagger
	}

	// 加载中间件
	routers.Use(handlers...)

	// 构建路由
	buildRouters(routers)

	routers.NoRoute(func(ctx *gin.Context) {
		response.GetResponse().SetHttpCode(http.StatusNotFound).FailCode(ctx, http.StatusNotFound)
	})

	return
}

// buildSwagger 创建swagger文档
func buildSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
