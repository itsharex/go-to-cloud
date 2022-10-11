package registry

import (
	"go-to-cloud/internal/repositories"
)

func ListRepositories(repoId uint) (repos []string, err error) {
	url, user, password, isSecurity, err := repositories.GetArtifactRepoByID(repoId)
	if err != nil {
		return
	}

	hub, err := GetRegistryHub(isSecurity, url, user, password)

	if err == nil {
		repos, err = hub.Repositories()
	}

	return
}
