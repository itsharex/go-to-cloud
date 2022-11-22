package agent

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/agent/client"
	gotocloud "go-to-cloud/internal/agent/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestGitClone(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	addr := "localhost:50010"
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"HealthCheckConfig": {"ServiceName": "%s"}}`, HEALTHCHECK_SERVICE)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(new(client.AccessTokenAuth)),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gotocloud.NewAgentClient(conn)

	// Contact the server and print out its response.
	r, err := c.GitClone(context.Background(), &gotocloud.CloneRequest{Address: "addr", Branch: "bra", EncodedToken: "token"})
	assert.NoError(t, err)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetWorkdir())
}

func TestShell(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	addr := "localhost:50010"
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"HealthCheckConfig": {"ServiceName": "%s"}}`, HEALTHCHECK_SERVICE)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(new(client.AccessTokenAuth)),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gotocloud.NewAgentClient(conn)

	s4 := &gotocloud.RunCommandRequest{
		Command: "echo",
		Args:    []string{"hello world"},
	}
	s3 := &gotocloud.RunCommandRequest{
		Command: "pwd",
		Next:    s4,
	}

	s2 := &gotocloud.RunCommandRequest{
		Command: "ls",
		Args:    []string{"."},
		Next:    s3,
	}

	s1 := &gotocloud.RunRequest{
		Workdir: "/",
		Command: &gotocloud.RunCommandRequest{
			Command: "cd",
			Args:    []string{"/"},
			Next:    s2,
		},
	}

	r, err := c.Run(context.Background(), s1)

	assert.NoError(t, err)
	assert.NotEmpty(t, r.Ret)
}
