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
		models[i] = deploy.Deployment{
			ProjectId:               deployments[i].ProjectId,
			K8sNamespace:            deployments[i].K8sNamespace,
			K8sRepoId:               deployments[i].K8sRepoId,
			ArtifactDockerImageId:   deployments[i].ArtifactDockerImageId,
			Ports:                   string(deployments[i].Ports),
			Cpus:                    deployments[i].Cpus,
			Env:                     string(deployments[i].Env),
			Replicas:                deployments[i].Replicas,
			Liveness:                deployments[i].Liveness,
			Readiness:               deployments[i].Readiness,
			RollingMaxSurge:         deployments[i].RollingMaxSurge,
			RollingMaxUnavailable:   deployments[i].RollingMaxUnavailable,
			ResourceLimitCpuRequest: deployments[i].ResourceLimitCpuRequest,
			ResourceLimitCpuLimits:  deployments[i].ResourceLimitCpuLimits,
			ResourceLimitMemRequest: deployments[i].ResourceLimitMemRequest,
			ResourceLimitMemLimits:  deployments[i].ResourceLimitMemLimits,
			NodeSelector:            string(deployments[i].NodeSelector),
		}
	}

	return models, nil
}
