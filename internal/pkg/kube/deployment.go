package kube

import (
	"context"
	"fmt"
	"go-to-cloud/internal/models/deploy"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

const DeploymentLabelSelector string = "gotocloud"

// GetDeployments 获取部署工作负载
func (client *Client) GetDeployments(ctx context.Context, ns string, projectId uint) ([]deploy.DeploymentDescription, error) {
	deployments, err := client.clientSet.AppsV1().Deployments(ns).List(ctx, metav1.ListOptions{
		LabelSelector: fmt.Sprintf("deployed=%s", DeploymentLabelSelector),
	})

	if err != nil {
		return nil, err
	}

	rlt := make([]deploy.DeploymentDescription, len(deployments.Items))

	for i, item := range deployments.Items {
		rlt[i] = deploy.DeploymentDescription{
			Id: func() uint {

			},
			Replicate:       0,
			AvailablePods:   0,
			UnavailablePods: 0,
			CreatedAt:       time.Time{},
			Conditions:      nil,
		}
	}
	return nil
}
