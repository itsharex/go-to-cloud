package project

import (
	"encoding/json"
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
			Namespace: deployments[i].K8sNamespace,
			K8S:       deployments[i].K8sRepoId,
			Artifact:  deployments[i].ArtifactDockerImageId,
			Ports: func() []struct {
				ServicePort   string `json:"text"`
				ContainerPort string `json:"value"`
			} {
				var t []struct {
					ServicePort   string `json:"text"`
					ContainerPort string `json:"value"`
				}
				if json.Unmarshal(deployments[i].Ports, &t) != nil {
					return nil
				} else {
					return t
				}
			}(),
			Env: func() []struct {
				VarName  string `json:"varName"`
				VarValue string `json:"varValue"`
			} {
				var t []struct {
					VarName  string `json:"varName"`
					VarValue string `json:"varValue"`
				}
				if json.Unmarshal(deployments[i].Env, &t) != nil {
					return nil
				} else {
					return t
				}
			}(),
			Replicate:  deployments[i].Replicas,
			CpuRequest: deployments[i].ResourceLimitCpuRequest,
			CpuLimits:  deployments[i].ResourceLimitCpuLimits,
			MemRequest: deployments[i].ResourceLimitMemRequest,
			MemLimits:  deployments[i].ResourceLimitMemLimits,
		}
	}

	return models, nil
}
