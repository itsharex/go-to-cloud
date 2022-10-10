package kube

import (
	"io"
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

var agentApplyCfg = &AppDeployConfig{
	Name: "go-to-cloud-agent",
	Ports: []Port{
		{
			ServicePort:   8080,
			ContainerPort: 8080,
			NodePort:      31080,
			PortName:      "go-to-cloud-agent",
		},
	},
	PortType: "NodePort",
	Image:    "nginx:latest",
	Replicas: 1,
}

func TestSetupAgentPodYaml(t *testing.T) {
	tpl1, err := template.New("testing").Parse(YamlTplService)
	assert.NoError(t, err)
	assert.NotNil(t, tpl1)
	err = tpl1.Execute(os.Stdout, agentApplyCfg)
	assert.NoError(t, err)

	tpl2, err := template.New("testing").Parse(YamlTplDeployment)
	assert.NoError(t, err)
	assert.NotNil(t, tpl2)
	err = tpl1.Execute(os.Stdout, agentApplyCfg)
	assert.NoError(t, err)
}

func TestSetupAgentPod(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to depend k8s")
	}

	ns := "go-to-cloud"

	file, err := os.Open("setup_test.yml")
	defer file.Close()
	k8scfgbyte, err := io.ReadAll(file)
	assert.NoError(t, err)
	k8scfg := string(k8scfgbyte)

	assert.NoError(t, Apply(&k8scfg, &ns, agentApplyCfg))
}