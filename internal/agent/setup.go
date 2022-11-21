package agent

import (
	"errors"
	"go-to-cloud/conf"
	"go-to-cloud/internal/agent/vars"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

// Setup 安装指定组织的agent
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
		Name:      vars.AgentNodeName,
		Ports: []kube.Port{
			{
				ContainerPort: 80,
				ServicePort:   80,
				NodePort:      agent.NodePort,
			},
		},
		Replicas: 1,
		Image:    *conf.GetAgentImage(),
	}

	client, err := kube.NewClient(&agent.KubeConfig)
	if err != nil {
		return err
	}

	return client.Launch(deploy)
}
