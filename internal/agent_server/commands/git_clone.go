package commands

import (
	agent "go-to-cloud/internal/agent_server"
	gotocloud "go-to-cloud/internal/agent_server/proto"
	"go-to-cloud/internal/utils"
)

func GitClone(src, branch, accessToken string) error {
	cmd := &gotocloud.RunRequest{
		Workdir: "./build",
		WorkId:  8,
		Command: &gotocloud.RunCommandRequest{
			Command: "git",
			Args: []string{
				src,
				branch,
				utils.Base64AesEny([]byte(accessToken)),
			},
		},
	}

	return agent.Runner.Execute(cmd)
}
