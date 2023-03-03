package monitor

import (
	"errors"
	"fmt"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

func ScaleApps(k8sRepoId uint, scale *deploy.ScalePods) error {
	k8sRepo, err := repositories.QueryK8sRepoById(k8sRepoId)
	if err != nil {
		return err
	}
	if k8sRepo == nil {
		return errors.New("部署环境丢失")
	}

	deployment, err := repositories.GetDeploymentById(scale.Id)
	if err != nil {
		return err
	}
	if deployment == nil {
		return errors.New("应用部署信息丢失")
	}

	ns := deployment.K8sNamespace
	deploymentName := fmt.Sprintf("%s-deployment", deployment.ArtifactDockerImageRepo.Name)

	k8sClient, err := kube.NewClient(&k8sRepo.KubeConfig)
	if err != nil {
		return err
	}

	return k8sClient.Scale(&ns, &deploymentName, scale.Num)
}
