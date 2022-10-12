package main

import (
	"flag"
	"go-to-cloud/internal/routers"
	"os"
	"strings"
)

var aType = flag.String("type", "", "运行方式: agent / web")
var aPort = flag.String("port", "", "端口")

// runMode 获取运行方式
// @bool: true:web; false: agent
// @string: 端口
func runMode() (bool, string) {
	// 优先读取命令行参数，其次使用go env，最后使用默认值
	flag.Parse()

	if len(*aType) == 0 {
		*aType = os.Getenv("type")
	}

	if len(*aPort) == 0 {
		*aPort = os.Getenv("port")
	}

	if len(*aType) == 0 {
		*aType = "web"
	}

	if len(*aPort) == 0 {
		*aPort = ":80"
	}

	if !strings.HasPrefix(*aPort, ":") {
		*aPort = ":" + *aPort
	}

	return strings.EqualFold(*aType, "web"), *aPort
}

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @BasePath /api
func main() {
	runType, port := runMode()
	if runType {
		// web模式运行
		_ = routers.SetRouters().Run(port)
	} else {
		// TODO: k8s agent模式运行
		// _ = agent.Startup()
	}
}
