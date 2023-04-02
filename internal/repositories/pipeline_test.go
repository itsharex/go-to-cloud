package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/internal/models/pipeline"
	"testing"
)

func TestNewPlan(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	mockstr := "mock"
	mockint := uint(1)
	assert.NoError(t, NewPlan(uint(0), uint(1), &pipeline.PlanModel{
		Name:            "name",
		QaEnabled:       true,
		ArtifactEnabled: false,
		UnitTest:        &mockstr,
		LintCheck:       &mockstr,
		Dockerfile:      &mockstr,
		ArtifactRepoId:  &mockint,
		Remark:          "remark",
		SourceCodeID:    1,
		Env:             "3.1",
		Branch:          "m",
	}))
}
