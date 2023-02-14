package builders

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"text/template"
)

func TestBuildPodSpec(t *testing.T) {

	tpl, err := template.New("k8s").Parse(BuildTemplate)
	assert.NoError(t, err)

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, PodSpecConfig{
		TaskName:   "Test",
		SourceCode: "DFDF.git",
		Sdk:        "6.0",
		Steps: []Step{
			{
				Command: "c1",
			},
			{
				Command: "c2",
			},
		},
	})

	assert.NoError(t, err)

	yml := strings.TrimSpace(buf.String())

	assert.True(t, len(yml) > 0)
}
