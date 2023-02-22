package project

import (
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/repositories"
)

func ListDeployments(projectId uint) ([]deploy.Deployment, error) {
	deployments, err := repositories.QueryDeploymentsByProjectId(projectId)

	if err != nil {
		return nil, err
	}

	models := make([]deploy.Deployment, len(deployments))
	for i := range deployments {
		models[i] = deploy.Deployment{Deployment: deployments[i]}
	}

	return models, nil
}
