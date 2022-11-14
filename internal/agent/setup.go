package agent

import (
	"errors"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

// Setup 安装agent至指定组织
func Setup(orgID uint) error {
	// 读取配置
	agent, err := repositories.GetAgentByOrgId(orgID)

	if err != nil {
		return err
	}

	if agent == nil {
		return errors.New("没有找到agent配置")
	}

	deploy := &kube.AppDeployConfig{
		Namespace: agent.Namespace,
		Name:      agentNodeName,
		Ports: []kube.Port{
			{
				ContainerPort: 80,
				ServicePort:   80,
				NodePort:      agent.NodePort,
			},
		},
		Replicas: 1,
		Image:    "-",
	}

	client, err := kube.NewClient(&agent.KubeConfig)
	if err != nil {
		return err
	}

	return client.Launch(deploy)
}
