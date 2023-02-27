package kube

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Scale 伸缩Pods数量
func (client *Client) Scale(ns, deployment *string, num int32) error {
	deploy, err := client.clientSet.AppsV1().Deployments(*ns).Get(context.TODO(), *deployment, metav1.GetOptions{})
	if err != nil {
		return err
	}
	*deploy.Spec.Replicas = num
	if _, err := client.clientSet.AppsV1().Deployments(*ns).Update(context.TODO(), deploy, metav1.UpdateOptions{}); err != nil {
		return err
	}

	return nil
}
