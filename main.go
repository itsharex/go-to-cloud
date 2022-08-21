package main

import (
	"flag"
	"go-to-cloud/internal/agent"
	"go-to-cloud/internal/routers"
	"strings"
)

var aType = flag.String("type", "agent", "运行方式: agent / web")

func main() {
	flag.Parse()

	if strings.EqualFold("web", *aType) {
		// web模式运行
		_ = routers.SetRouters().Run(":8080")
	} else {
		// k8s agent模式运行
		_ = agent.Startup()
	}
}
