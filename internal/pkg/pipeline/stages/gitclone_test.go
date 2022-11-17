package stages

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGithubCloneFromPublicRepo(t *testing.T) {
	stage := GitCloneStage{
		GitUrl:  "https://github.com/go-git/go-git.git",
		WorkDir: "",
		Token:   "",
		Branch:  "pr-1152",
	}
	err := stage.Run()

	assert.NoError(t, err)
}

func TestGithubCloneFromPrivateRepo(t *testing.T) {
	if testing.Short() {
		t.Skip("token expired, try another valid token")
	}
	stage := GitCloneStage{
		GitUrl:  "https://github.com/go-to-cloud/go-to-cloud.git",
		WorkDir: "",
		Token:   "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",
	}
	err := stage.Run()

	assert.NoError(t, err)
}

func TestGithubCloneFromPrivateRepoToLocal(t *testing.T) {
	if testing.Short() {
		t.Skip("token expired, try another valid token")
	}
	dir, err := os.MkdirTemp("", "workspace")
	stage := GitCloneStage{
		GitUrl:  "https://github.com/go-to-cloud/go-to-cloud.git",
		Token:   "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",
		WorkDir: dir,
	}
	defer os.RemoveAll(dir)

	err = stage.Run()

	assert.NoError(t, err)
}
