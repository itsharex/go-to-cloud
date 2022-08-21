package agent

import (
	"fmt"
	"go-to-cloud/internal/pkg/kube"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"

	appsv1 "k8s.io/client-go/applyconfigurations/apps/v1"
)

var app *kube.AppDeployConfig

func initAgentConfig(workPort, nodePort int, image *string) *kube.AppDeployConfig {
	return &kube.AppDeployConfig{
		Name: "go-to-cloud-agent",
		Ports: []kube.Port{
			{
				ServicePort:   workPort,
				ContainerPort: workPort,
				NodePort:      nodePort,
				PortName:      "agent",
			},
		},
		PortType: "NodePort",
		Image:    *image,
		Replicas: 1,
	}
}

// Setup 将Agent安装至k8s常驻
func Setup(namespace, agentImage, k8sConfig *string, workPort, nodePort int) (bool, error) {
	client, err := kube.NewClient(k8sConfig)

	_, err = client.GetOrAddNamespace(namespace)
	if err != nil {
		panic(err)
	}

	deploy, service, err := kube.GetYamlFromTemple(initAgentConfig(workPort, nodePort, agentImage))

	deployCfg := appsv1.DeploymentApplyConfiguration{}

	if err := kube.DecodeYaml(deploy, &deployCfg); err != nil {
		fmt.Println(err)
		return false, err
	}

	serviceCfg := corev1.ServiceApplyConfiguration{}
	if err := kube.DecodeYaml(service, &serviceCfg); err != nil {
		fmt.Println(err)
		return false, err
	}

	if _, e := client.ApplyDeployment(namespace, &deployCfg); e != nil {
		fmt.Println(e)
		return false, err
	}

	if _, e := client.ApplyService(namespace, &serviceCfg); e != nil {
		fmt.Println(e)
		return false, err
	}

	return true, nil
}
