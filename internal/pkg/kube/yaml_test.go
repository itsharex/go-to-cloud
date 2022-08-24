package kube

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"text/template"
)

func TestYamlTemplateParser(t *testing.T) {

	cpuRequest, cpuLimit, memRequest, memLimit := 1, 2, 3, 4
	config := AppDeployConfig{
		Name:     "HelloWorld",
		Image:    "Empty",
		Replicas: 1,
		Ports: []Port{
			{
				ServicePort:   10,
				ContainerPort: 10,
			},
			{
				ServicePort:   11,
				ContainerPort: 11,
				NodePort:      91,
			},
		},
		ResourceLimit: &ResLimits{
			CpuRequest: &cpuRequest,
			CpuLimits:  &cpuLimit,
			MemRequest: &memRequest,
			MemLimits:  &memLimit,
		},
		RollingUpdate: &RollingUpdateStrategy{
			MaxSurge:       12,
			MaxUnavailable: 30,
		},
		Dependencies: []DependContainer{
			{
				ContainerName: "C1",
				Namespace:     "helloWorld",
			},
			{
				ContainerName: "c2",
				Namespace:     "helloWorld2",
			},
		},
	}

	assert.Error(t, config.validate())

	tpl, err := template.New("deploy").Parse(YamlTplDeployment)
	assert.NoError(t, err)

	tpl, err = template.New("service").Parse(YamlTplService)
	assert.NoError(t, err)

	assert.Equal(t, "NodePort", config.PortType)
	err = tpl.Execute(os.Stdout, config)

	assert.NoError(t, err)
}

func TestYamlTemplateCheck(t *testing.T) {

	cpuRequest, cpuLimit, memRequest, memLimit := 1, 2, 3, 4

	config := AppDeployConfig{
		Name:     "HelloWorld",
		Image:    "Nginx:latest",
		Replicas: 1,
		ResourceLimit: &ResLimits{
			CpuRequest: &cpuRequest,
			CpuLimits:  &cpuLimit,
			MemRequest: &memRequest,
			MemLimits:  &memLimit,
		},
		RollingUpdate: &RollingUpdateStrategy{
			MaxSurge:       12,
			MaxUnavailable: 30,
		},
		Dependencies: []DependContainer{
			{
				ContainerName: "C1",
				Namespace:     "helloWorld",
			},
			{
				ContainerName: "c2",
				Namespace:     "helloWorld2",
			},
		},
	}

	assert.NoError(t, config.validate())
}
