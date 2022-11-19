package stages

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSh(t *testing.T) {
	o, e := sh.NewSession().SetInput("hl").Command("echo", "cat").Output()
	assert.NoError(t, e)
	fmt.Println(string(o))
}

func TestShell_Run(t *testing.T) {

	c4 := ShellCommand{
		Command: "echo",
		Args:    []string{"Hello World!"},
	}
	c3 := ShellCommand{
		Command: "pwd",
		Next:    &c4,
	}
	c2 := ShellCommand{
		Command: "ls",
		Args:    []string{"."},
		Next:    &c3,
	}
	c1 := ShellCommand{
		Command: "cd",
		Args:    []string{"/"},
		Next:    &c2,
	}
	shell := &Shell{
		Commands: c1,
		WorkDir:  "",
	}

	assert.NoError(t, shell.Run())
	fmt.Println(shell.Result)
}
