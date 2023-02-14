package lang

import (
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/repositories"
)

type Golang struct {
}

func (m *Golang) Sdk(env string) string {
	const dotnet = "golang"
	switch env {
	case Go117:
		return dotnet + ":1.17.0"
	case Go116:
		return dotnet + ":1.16.0"
	case Go120:
		return dotnet + ":1.20.0"
	case Go119:
		return dotnet + ":1.19.0"
	case Go118:
		return dotnet + ":1.18.0"
	}

	return dotnet + ":latest"
}

func (m *Golang) Steps(env string, steps []repositories.PipelineSteps) map[pipeline.PlanStepType]string {
	// TODO: steps
	panic("not implemented")
}
