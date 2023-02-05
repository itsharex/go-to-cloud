package commands

import (
	"github.com/stretchr/testify/assert"
	server "go-to-cloud/internal/agent_server"
	"testing"
	"time"
)

func TestGitClone(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	go func() {
		tick := time.Tick(time.Second * 10)
		<-tick
		assert.NoError(t, GitClone("src", "branch", "token"))
	}()

	port := "50010"
	server.Startup(&port)
}
