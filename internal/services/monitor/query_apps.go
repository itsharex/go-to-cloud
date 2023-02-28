package monitor

import (
	"errors"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

func QueryApps(projectId, k8sId uint) ([]deploy.DeploymentDescription, error) {
	repo, err1 := repositories.QueryK8sRepoById(k8sId)
	if err1 != nil {
		return nil, err1
	}

	client, err2 := kube.NewClient(&repo.KubeConfig)
	_ = client
	if err2 != nil {
		return nil, err2
	}

	// todo

	return nil, errors.New("not implement")
}
