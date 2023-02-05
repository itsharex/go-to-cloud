package commands

import (
	agent "go-to-cloud/internal/agent_server"
	gotocloud "go-to-cloud/internal/agent_server/proto"
	"go-to-cloud/internal/utils"
)

func GitClone(src, branch, accessToken string) error {
	subCmd := &gotocloud.RunCommandRequest{
		Command: "dotnet build",
	}

	cmd := &gotocloud.RunRequest{
		Workdir: "./build",
		Command: &gotocloud.RunCommandRequest{
			Command: "git",
			Args: []string{
				src,
				branch,
				utils.Base64AesEny([]byte(accessToken)),
			},
			Next: subCmd,
		},
	}

	return agent.Runner.Execute(cmd)
}
