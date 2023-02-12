package project

import (
	"github.com/stretchr/testify/assert"
	gotocloud "go-to-cloud/internal/agent_server/proto"
	"go-to-cloud/internal/repositories"
	"testing"
)

func TestMakeRequestCommand(t *testing.T) {
	plan := &repositories.Pipeline{
		SourceCode: repositories.ProjectSourceCode{
			GitUrl:   "giturl",
			CodeRepo: repositories.CodeRepo{AccessToken: "accesstoken"},
		},
		Branch: "branch",
		PipelineSteps: []repositories.PipelineSteps{
			{},
			{},
		},
	}

	req := &gotocloud.RunRequest{WorkId: 1}

	req.Command = makeGitCloneRequestCommand(plan)

	assert.Equal(t, 3, len(req.Command.Args))
	assert.Equal(t, "git", req.Command.Command)

	command := req.Command
	for _, step := range plan.PipelineSteps {
		makeShellRequestCommand(command, &step)
		command = command.Next
	}

	assert.True(t, nil == req.Command.Next.Next.Next)
}
