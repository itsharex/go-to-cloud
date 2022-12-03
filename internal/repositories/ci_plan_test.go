package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/models/build"
	"testing"
)

func TestNewPlan(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	assert.NoError(t, NewPlan(uint(0), uint(1), &build.PlanModel{
		Name:            "name",
		QaEnabled:       true,
		ArtifactEnabled: false,
		UnitTest:        "test",
		LintCheck:       "lintcheck",
		Dockerfile:      "dockerfile",
		ArtifactRepoId:  1,
		Remark:          "remark",
		SourceCodeId:    1,
		BuildEnv:        "3.1",
		Branch:          "branch",
	}))
}
