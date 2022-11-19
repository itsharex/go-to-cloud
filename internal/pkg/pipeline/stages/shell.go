package stages

import (
	"errors"
)
import "github.com/codeskyblue/go-sh"

type ShellCommand struct {
	Command string        `json:"command"`
	Args    []string      `json:"args"`
	Next    *ShellCommand `json:"next"`
}
type Shell struct {
	Commands ShellCommand `json:"commands"`
	WorkDir  string       `json:"workDir"`
	Result   string       `json:"result"`
}

func (m *Shell) Stub() error {

	// TODO: master调用
	return errors.New("NOT Implemented")
}

func (m *Shell) Run() error {
	session := sh.NewSession()
	if !individualCmd(session, &m.Commands) {
		session = session.Command(m.Commands.Command, m.Commands.Args)
	}
	next := m.Commands.Next
	for next != nil {
		if !individualCmd(session, next) {
			session = session.Command(next.Command, next.Args)
		}
		next = next.Next
	}
	r, err := session.Output()
	if err == nil {
		m.Result = string(r)
	}
	return err
}

func individualCmd(session *sh.Session, command *ShellCommand) bool {
	if command.Command == "cd" {
		if len(command.Args) > 0 {
			session.SetDir(command.Args[0])
		}
		return true
	}

	return false
}
