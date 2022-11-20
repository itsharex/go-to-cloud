package server

import (
	gotocloud "go-to-cloud/internal/agent/proto"
)

type Agent struct {
	gotocloud.UnimplementedAgentServer
}
