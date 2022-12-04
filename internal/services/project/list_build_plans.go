package project

import (
	"go-to-cloud/internal/models/build"
	"go-to-cloud/internal/repositories"
)

func ListBuildPlans(projectId uint) ([]build.PlanCardModel, error) {
	plans, err := repositories.QueryPlan(projectId)

	if err != nil {
		return nil, err
	}

	models := make([]build.PlanCardModel, len(plans))
	for i, plan := range plans {
		unitTestEnabled, lintEnabled, artifactEnabled := false, false, false
		var unitTest, lint *string = nil, nil
		for _, step := range plan.CiPlanSteps {
			if step.Type == 1 {
				unitTestEnabled = true
				unitTest = &step.Script
				continue
			}
			if step.Type == 2 {
				lintEnabled = true
				lint = &step.Script
				continue
			}
			if step.Type == 4 {
				artifactEnabled = true
				continue
			}
		}
		models[i] = build.PlanCardModel{
			PlanModel: build.PlanModel{
				Id:              plan.ID,
				Name:            plan.Name,
				Env:             plan.Env,
				SourceCodeID:    plan.SourceCodeID,
				Branch:          plan.Branch,
				QaEnabled:       unitTestEnabled || lintEnabled,
				UnitTest:        unitTest,
				LintCheck:       lint,
				ArtifactEnabled: artifactEnabled,
			},
			LastBuildAt:     plan.LastBuildAt,
			LastBuildResult: plan.LastBuildResult,
		}
	}

	return models, nil
}
