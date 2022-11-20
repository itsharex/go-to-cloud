# GO-TO-CLOUD

[![Build Status](https://github.com/go-to-cloud/go-to-cloud/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/go-to-cloud/go-to-cloud/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/go-to-cloud/go-to-cloud/branch/main/graph/badge.svg?token=9Y81AN6KUA)](https://codecov.io/gh/go-to-cloud/go-to-cloud)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/go-to-cloud/go-to-cloud/blob/main/LICENSE)

### Master 后端服务


### Swagger

#### 使用swaggo管理接口

- 首次使用需要安装swag
```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

- 创建（或更新）swagger文档

```shell
swag init
```

- 访问 [swagger api](http://localhost:8080/swagger/index.html)

#### 接口文档编写参考

```
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context)  {
	g.JSON(http.StatusOK, "helloworld")
}
```

# Agent

> Agent 运行在构建节点机器上，与Master使用gPrc通讯

1. 下载[protoc]("https://github.com/protocolbuffers/protobuf/releases")
> 下载完成后建议将`protoc`路径配置到PATH

2. 安装gRPC工具

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

3. 生成pb文件

```shell
protoc ./internal/agent/proto/*.proto --go_out=./internal/agent/proto --go-grpc_out=./internal/agent/proto
```