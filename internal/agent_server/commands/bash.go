package commands

import (
	agent "go-to-cloud/internal/agent_server"
	gotocloud "go-to-cloud/internal/agent_server/proto"
)

func Bash(sh *gotocloud.RunRequest) error {
	return agent.Runner.Execute(sh)
}
