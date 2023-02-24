package project

import (
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/repositories"
	"gorm.io/datatypes"
)

func CreateDeployments(_ uint64, d *deploy.Deployment) error {
	repo := repositories.Deployment{
		ProjectId:               d.ProjectId,
		K8sNamespace:            d.K8sNamespace,
		K8sRepoId:               d.K8sRepoId,
		ArtifactDockerImageId:   d.ArtifactDockerImageId,
		Ports:                   datatypes.JSON(d.Ports),
		Cpus:                    d.Cpus,
		Env:                     datatypes.JSON(d.Env),
		Replicas:                d.Replicas,
		Liveness:                d.Liveness,
		Readiness:               d.Readiness,
		RollingMaxSurge:         d.RollingMaxSurge,
		RollingMaxUnavailable:   d.RollingMaxUnavailable,
		ResourceLimitCpuRequest: d.ResourceLimitCpuRequest,
		ResourceLimitCpuLimits:  d.ResourceLimitCpuLimits,
		ResourceLimitMemRequest: d.ResourceLimitMemRequest,
		ResourceLimitMemLimits:  d.ResourceLimitMemLimits,
		NodeSelector:            datatypes.JSON(d.NodeSelector),
	}
	return repositories.CreateDeployment(&repo)
}
