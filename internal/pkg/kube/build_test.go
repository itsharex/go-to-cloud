package kube

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/models/pipeline"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	podSpecConfig := &PodSpecConfig{
		Namespace:  "testNs",
		TaskName:   "testTask",
		SourceCode: "testSource",
		Sha:        "test-branch",
		Sdk:        "testSdk",
		Steps: []Step{
			{
				CommandType: pipeline.Image,
				CommandText: "image",
				Command:     "cli 1",
				Dockerfile:  "dockerfile.file",
				Registry: struct {
					Url      string
					User     string
					Password string
					Security bool
				}{Url: "regUrl", User: "regUser", Password: "regPwd", Security: false},
			},
			{
				CommandType: pipeline.LintCheck,
				CommandText: "image",
				Command:     "cli 2",
			},
		},
	}

	spec, err := makeTemplate(podSpecConfig)
	assert.NoError(t, err)

	err = DecodeYaml(spec, &corev1.PodApplyConfiguration{})
	fmt.Println(*spec)
	assert.NoError(t, err)
}
