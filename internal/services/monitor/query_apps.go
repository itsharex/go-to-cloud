package monitor

import (
	"context"
	"errors"
	"fmt"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// QueryApps 查找应用
// projectId 如果为0表示查找所有项目的所有应用
// deploymentId 如果为0表示查找所有应用
func QueryApps(projectId, deploymentId, k8sId uint) ([]deploy.DeploymentDescription, error) {
	repo, err := repositories.QueryK8sRepoById(k8sId)
	if err != nil {
		return nil, err
	}

	client, err := kube.NewClient(&repo.KubeConfig)
	if err != nil {
		return nil, err
	}

	deployments, err := func() ([]repositories.Deployment, error) {
		if a, err := repositories.QueryDeploymentsByK8s(k8sId); err != nil {
			return nil, err
		} else if projectId == 0 {
			return a, nil
		} else {
			f := make([]repositories.Deployment, 0)
			for i, deployment := range a {
				if deployment.ProjectId == projectId {
					f = append(f, a[i])
				}
			}
			return f, nil
		}
	}()

	if err != nil {
		return nil, err
	}

	nsAll := make([]string, 0)

	var deploymentIdFilter *map[uint]bool
	if deploymentId > 0 {
		m := make(map[uint]bool)
		deploymentIdFilter = &m
	}
	for _, deployment := range deployments {
		nsAll = append(nsAll, deployment.K8sNamespace)
		if deploymentIdFilter != nil {
			(*deploymentIdFilter)[deployment.ID] = true
		}
	}
	ns := utils.Distinct(nsAll)

	for _, namespace := range ns {
		client.GetDeployments(context.TODO(), namespace, deploymentIdFilter)
	}

	return nil, errors.New("not implement")
}

func ab() {
	a, err := returnErr("a")
	_ = a
	fmt.Println(err.Error())

	b, err := returnErr("b")
	_ = b
	fmt.Println(err.Error())
}

func returnErr(str string) (string, error) {
	return str, errors.New(str)
}
