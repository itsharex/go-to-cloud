package kube

import (
	"fmt"
	"golang.org/x/net/context"
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	applyApps "k8s.io/client-go/applyconfigurations/apps/v1"
	applyCore "k8s.io/client-go/applyconfigurations/core/v1"
)

// ApplyDeployment kubectl apply -f yml
// Kind: Deployment
func (client *Client) ApplyDeployment(namespace *string, yml *applyApps.DeploymentApplyConfiguration) (*apps.Deployment, error) {
	client.getOrCreateNamespace(namespace)
	return client.clientSet.AppsV1().Deployments(*namespace).Apply(context.TODO(), yml, *client.defaultApplyOptions)
}

// ApplyService kubectl apply -f yml
// Kind: Service
func (client *Client) ApplyService(namespace *string, yml *applyCore.ServiceApplyConfiguration) (*core.Service, error) {
	client.getOrCreateNamespace(namespace)
	return client.clientSet.CoreV1().Services(*namespace).Apply(context.TODO(), yml, *client.defaultApplyOptions)
}

// Setup 安装Agent
func Setup(k8sConfig, namespace *string, pod *AppDeployConfig) error {
	client, err := NewClient(k8sConfig)

	_, err = client.GetOrAddNamespace(namespace)
	if err != nil {
		panic(err)
	}

	deploy, service, err := GetYamlFromTemple(pod)
	if err != nil {
		panic(err)
	}

	deployConfig := applyApps.DeploymentApplyConfiguration{}
	if err = DecodeYaml(deploy, &deployConfig); err != nil {
		fmt.Println(err)
		panic(err)

	}
	serviceConfig := applyCore.ServiceApplyConfiguration{}
	if err = DecodeYaml(service, &serviceConfig); err != nil {
		fmt.Println(err)
		panic(err)
	}

	if _, err = client.ApplyDeployment(namespace, &deployConfig); err != nil {
		fmt.Println(err)
		panic(err)
	}

	if _, err = client.ApplyService(namespace, &serviceConfig); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return nil
}
