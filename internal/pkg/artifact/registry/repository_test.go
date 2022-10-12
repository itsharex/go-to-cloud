package registry

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/repositories"
	"testing"
)

func TestListRepositoriesShouldSuccess(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repoID := uint(2)

	r, err := ListRepositories(repoID)

	assert.NotNil(t, r)
	assert.NoError(t, err)

	url, user, password, isSecurity, err := repositories.GetArtifactRepoByID(repoID)
	if err != nil {
		return
	}

	hub, err := GetRegistryHub(isSecurity, url, user, password)
	for _, s := range r {
		//reader, _ := hub.DownloadBlob(s, "sha256:1e45b6ec54a62f828a4066ddc934fff9467911fb0b0ebbb4660bd36340eb4a86")
		//buf := new(bytes.Buffer)
		//buf.ReadFrom(reader)
		//reader.Close()
		//newStr := buf.String()
		//fmt.Println(newStr)
		tags, _ := hub.Tags(s.FullName)
		for _, tag := range tags {
			fmt.Println(tag)
			//	m, _ := hub.ManifestV2(s, tag)
			//	reader, _ := hub.DownloadBlob(s, "sha256:1e45b6ec54a62f828a4066ddc934fff9467911fb0b0ebbb4660bd36340eb4a86")
			//	buf := new(bytes.Buffer)
			//	buf.ReadFrom(reader)
			//	reader.Close()
			//	newStr := buf.String()
			//	fmt.Println(newStr)
		}
	}
}
