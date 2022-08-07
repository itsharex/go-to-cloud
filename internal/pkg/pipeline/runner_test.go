package pipeline

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"text/template"
)

func TestRunnerTemplate(t *testing.T) {

	container := Container{
		Image: "mcr.microsoft.com/dotnet/sdk",
		Tag:   "6.0.0",
		Alias: "compile",
	}

	pod := IntegrationPod{
		Containers: []Container{
			container,
		},
		ImageName:  "tmp",
		ImageTag:   "latest",
		Dockerfile: "./Dockerfile",
		Workspace:  "./",
		Registry:   "fanhousanbu-docker.pkg.coding.net/gotocloud-artifacts/kaniko",
	}

	tpl, err := template.New("runner").Parse(integrationPodTemplate)
	assert.NoError(t, err)

	err = tpl.Execute(os.Stdout, pod)
	assert.NoError(t, err)
}
