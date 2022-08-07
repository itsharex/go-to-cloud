package stages

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGithubCloneFromPublicRepo(t *testing.T) {
	gitUrl := "https://github.com/go-git/go-git.git"
	dest := ""
	token := ""
	err := gitClone(&gitUrl, &dest, &token)

	assert.NoError(t, err)
}

func TestGithubCloneFromPrivateRepo(t *testing.T) {
	if testing.Short() {
		t.Skip("token expired, try another valid token")
	}
	gitUrl := "https://github.com/go-to-cloud/go-to-cloud.git"
	dest := ""
	token := "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"
	err := gitClone(&gitUrl, &dest, &token)

	assert.NoError(t, err)
}

func TestGithubCloneFromPrivateRepoToLocal(t *testing.T) {
	if testing.Short() {
		t.Skip("token expired, try another valid token")
	}
	gitUrl := "https://github.com/go-to-cloud/go-to-cloud.git"
	dir, err := os.MkdirTemp("", "workspace")
	defer os.RemoveAll(dir)
	token := "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"
	err = gitClone(&gitUrl, &dir, &token)

	assert.NoError(t, err)
}
