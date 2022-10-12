package registry

import (
	registryModels "go-to-cloud/internal/models/artifact/registry"
	"go-to-cloud/internal/repositories"
	"strings"
)

func extractNameFromRepo(fullName *string) *string {
	split := strings.LastIndex(*fullName, "/")
	if split > 0 {
		name := (*fullName)[split:]
		return &name
	} else {
		return nil
	}
}

func ListRepositories(repoId uint) (images []registryModels.Image, err error) {
	url, user, password, isSecurity, err := repositories.GetArtifactRepoByID(repoId)
	if err != nil {
		return
	}

	hub, err := GetRegistryHub(isSecurity, url, user, password)

	if err == nil {
		repos, err := hub.Repositories()

		if err != nil {
			return nil, err
		}

		images = make([]registryModels.Image, len(repos))
		for i, repo := range repos {
			if tags, err := hub.Tags(repo); err == nil {
				images[i] = registryModels.Image{
					Name:     *extractNameFromRepo(&repo),
					FullName: repo,
					Tags:     tags,
				}
			}
		}
	}

	return
}
