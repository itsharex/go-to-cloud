package main

import (
	"flag"
	"fmt"
	"go-to-cloud/conf"
	"go-to-cloud/internal/routers"
	"strings"
)

var aType = flag.String("type", "agent", "运行方式: agent / web")
var aPort = flag.String("port", ":8080", "端口")

var confFile *string

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func main() {
	flag.Parse()

	if strings.EqualFold("web", *aType) {
		if len(*aPort) == 0 {
			*aPort = ":8080"
		}
		// web模式运行
		_ = routers.SetRouters().Run(*aPort)
	} else {
		// TODO: k8s agent模式运行
		// _ = agent.Startup()
	}
}

func init() {
	envName := conf.Environment.GetEnvName()

	var tmp string
	tmp, confFile = fmt.Sprintf("conf/appsettings.%s.yaml", *envName), &tmp
}
