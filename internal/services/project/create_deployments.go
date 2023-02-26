package project

import (
	"encoding/json"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/repositories"
	"gorm.io/datatypes"
)

func CreateDeployments(projectId uint, d *deploy.Deployment) error {
	ser := func(v any) string {
		if m, e := json.Marshal(v); e != nil {
			return ""
		} else {
			return string(m)
		}
	}
	v := func(l uint) uint {
		if d.EnableLimit {
			return l
		} else {
			return 0
		}
	}
	repo := repositories.Deployment{
		ProjectId:               projectId,
		K8sNamespace:            d.Namespace,
		K8sRepoId:               d.K8S,
		ArtifactDockerImageId:   d.Artifact,
		ArtifactTag:             d.ArtifactTag,
		Ports:                   datatypes.JSON(ser(d.Ports)),
		Env:                     datatypes.JSON(ser(d.Env)),
		Replicas:                d.Replicate,
		ResourceLimitCpuRequest: v(d.CpuRequest),
		ResourceLimitCpuLimits:  v(d.CpuLimits),
		ResourceLimitMemRequest: v(d.MemRequest),
		ResourceLimitMemLimits:  v(d.MemLimits),
	}
	return repositories.CreateDeployment(&repo)
}
