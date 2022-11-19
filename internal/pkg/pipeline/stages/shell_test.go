package stages

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShell_Run(t *testing.T) {
	shell := &Shell{
		Commands: ShellCommand{
			Command: "ls",
			Args:    []string{"."},
			Next:    nil,
		},
		WorkDir: "",
	}

	assert.NoError(t, shell.Run())
	fmt.Println(shell.Result)
}
