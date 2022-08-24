package main

import (
	"flag"
	"go-to-cloud/internal/routers"
	"strings"
)

var aType = flag.String("type", "agent", "运行方式: agent / web")
var aPort = flag.String("port", ":8080", "端口")

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
