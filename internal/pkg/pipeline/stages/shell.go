package stages

import "errors"
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
	session = session.Command(m.Commands.Command, m.Commands.Args)
	next := m.Commands.Next
	for next != nil {
		session = session.Command(next.Command, next.Args)
		next = m.Commands.Next
	}
	r, err := session.Output()
	if err != nil {
		m.Result = string(r)
	}
	return err
}
