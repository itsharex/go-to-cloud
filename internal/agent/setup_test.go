package agent

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/pkg/kube"
	"io"
	"os"
	"testing"
	"text/template"
)

var podApplyCfg = &kube.PodApplyConfig{
	Name: "go-to-cloud-agent",
	Containers: []kube.PodContainer{
		{
			Image: "nginx:latest", // "go-to-cloud-docker.pkg.coding.net/devops/repo/kaniko:latest"
			Name:  "agent",
			TTY:   true,
			Ports: []kube.PodPort{
				{
					Port:     80,
					HostPort: 10080,
				},
			},
		},
	},
}

func TestSetupAgentPodYaml(t *testing.T) {
	tpl, err := template.New("pod").Parse(kube.YamlTplPod)
	assert.NoError(t, err)
	assert.NotNil(t, tpl)
	err = tpl.Execute(os.Stdout, podApplyCfg)
	assert.NoError(t, err)
}

func TestSetupAgentPod(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to depend k8s")
	}

	ns := "go-to-cloud"

	file, err := os.Open("k8s.config.test.yml")
	defer file.Close()
	k8scfgbyte, err := io.ReadAll(file)
	assert.NoError(t, err)
	k8scfg := string(k8scfgbyte)

	pod, err := Setup(&k8scfg, &ns, podApplyCfg)
	assert.NoError(t, err)
	assert.NotNil(t, pod)
}
