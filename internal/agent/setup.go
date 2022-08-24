package agent

import (
	"fmt"
	"go-to-cloud/internal/pkg/kube"
	appv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	cfgcorev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// Setup 安装Pod
func Setup(k8sConfig, namespace *string, pod *kube.AppDeployConfig) error {
	client, err := kube.NewClient(k8sConfig)

	_, err = client.GetOrAddNamespace(namespace)
	if err != nil {
		panic(err)
	}

	deploy, service, err := kube.GetYamlFromTemple(pod)
	if err != nil {
		panic(err)
	}

	deployConfig := appv1.DeploymentApplyConfiguration{}
	if err = kube.DecodeYaml(deploy, &deployConfig); err != nil {
		fmt.Println(err)
		panic(err)

	}
	serviceConfig := cfgcorev1.ServiceApplyConfiguration{}
	if err = kube.DecodeYaml(service, &serviceConfig); err != nil {
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
