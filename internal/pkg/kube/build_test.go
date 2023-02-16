package kube

import (
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"testing"
)

func TestParseTemplate(t *testing.T) {
	podSpecConfig := &PodSpecConfig{
		Namespace:  "testNs",
		TaskName:   "testTask",
		SourceCode: "testSource",
		Sdk:        "testSdk",
		Steps: []Step{
			{
				CommandType: "type1",
				Command:     "cli 1",
			},
			{
				CommandType: "type2",
				Command:     "cli 2",
			},
		},
	}

	spec, err := makeTemplate(podSpecConfig)
	assert.NoError(t, err)

	err = DecodeYaml(spec, &corev1.PodApplyConfiguration{})
	assert.NoError(t, err)
}
