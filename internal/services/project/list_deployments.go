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
			Namespace:    deployments[i].K8sNamespace,
			K8S:          deployments[i].K8sRepoId,
			K8sName:      deployments[i].K8sRepo.Name,
			Artifact:     deployments[i].ArtifactDockerImageId,
			ArtifactName: deployments[i].ArtifactDockerImageRepo.Name,
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
				VarName  string `json:"text"`
				VarValue string `json:"value"`
			} {
				var t []struct {
					VarName  string `json:"text"`
					VarValue string `json:"value"`
				}
				t1 := make([]struct {
					VarName  string `json:"text"`
					VarValue string `json:"value"`
				}, 0)
				if json.Unmarshal(deployments[i].Env, &t) != nil {
					return nil
				} else {
					for i, s := range t {
						if len(s.VarName) > 0 {
							t1 = append(t1, t[i])
						}
					}
					return t1
				}
			}(),
			Replicate:   deployments[i].Replicas,
			CpuRequest:  deployments[i].ResourceLimitCpuRequest,
			CpuLimits:   deployments[i].ResourceLimitCpuLimits,
			MemRequest:  deployments[i].ResourceLimitMemRequest,
			MemLimits:   deployments[i].ResourceLimitMemLimits,
			ArtifactTag: deployments[i].ArtifactTag,
		}
	}

	return models, nil
}
