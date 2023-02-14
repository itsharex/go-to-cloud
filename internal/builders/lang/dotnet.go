package lang

import (
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/repositories"
)

type DotNet struct {
}

func (m *DotNet) Sdk(env string) string {
	const dotnet = "mcr.microsoft.com/dotnet/sdk"
	switch env {
	case DotNet6:
		return dotnet + ":6.0.0"
	case DotNet5:
		return dotnet + ":5.0.0"
	case DotNet7:
		return dotnet + ":7.0.0"
	case DotNet3:
		return dotnet + ":3.1.0"
	}

	return dotnet + ":latest"
}

func (m *DotNet) Steps(env string, steps []repositories.PipelineSteps) map[pipeline.PlanStepType]string {
	// TODO: steps
	panic("not implemented")
}
