package agent

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestSetupAgentPod(t *testing.T) {
	if testing.Short() {
		t.Skip("skipped due to depend k8s")
	}

	ns := "go-to-cloud"
	img := "go-to-cloud-docker.pkg.coding.net/devops/repo/kaniko:latest"

	file, err := os.Open("k8s.config.test.yml")
	defer file.Close()
	k8scfgbyte, err := io.ReadAll(file)
	assert.NoError(t, err)
	k8scfg := string(k8scfgbyte)

	Setup(&ns, &img, &k8scfg, 80, 31080)
}
