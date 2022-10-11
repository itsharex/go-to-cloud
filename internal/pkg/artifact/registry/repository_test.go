package registry

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/repositories"
	"testing"
)

func TestListRepositoriesShouldSuccess(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repoID := uint(1)

	r, err := ListRepositories(repoID)

	assert.NotNil(t, r)
	assert.NoError(t, err)

	url, user, password, isSecurity, err := repositories.GetArtifactRepoByID(repoID)
	if err != nil {
		return
	}

	hub, err := GetRegistryHub(isSecurity, url, user, password)
	for _, s := range r {
		tags, _ := hub.Tags(s)
		for _, tag := range tags {
			m, _ := hub.ManifestV2(s, tag)
			_ = m
		}
	}
}
