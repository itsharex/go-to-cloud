package server

import (
	"context"
	gotocloud "go-to-cloud/internal/agent/proto"
	"go-to-cloud/internal/pkg/pipeline/stages"
)

func (m *Agent) Run(ctx context.Context, in *gotocloud.RunRequest) (*gotocloud.RunResponse, error) {
	shell := stages.Shell{
		Commands: stages.ShellCommand{
			Command: in.Command.Command,
			Args:    in.Command.Args,
		},
		WorkDir: in.Workdir,
	}
	nextR := in.Command.Next
	var current *stages.ShellCommand
	current = &shell.Commands
	for nextR != nil {
		cmd := &stages.ShellCommand{
			Command: nextR.Command,
			Args:    nextR.Args,
		}
		(*current).Next = cmd
		current = cmd
		nextR = nextR.Next
	}

	err := shell.Run()
	if err != nil {
		return nil, err
	} else {
		return &gotocloud.RunResponse{
			Ret: shell.Result,
		}, nil
	}
}
