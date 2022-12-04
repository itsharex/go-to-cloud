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
		models[i] = build.PlanCardModel{
			PlanModel: build.PlanModel{
				Name:         plan.Name,
				Env:          plan.Env,
				SourceCodeID: plan.SourceCodeID,
				Branch:       plan.Branch,
			},
			LastBuildAt:     plan.LastBuildAt,
			LastBuildResult: plan.LastBuildResult,
		}
	}

	return models, nil
}
