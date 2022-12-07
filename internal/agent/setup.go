package agent

import (
	"errors"
	"go-to-cloud/conf"
	"go-to-cloud/internal/agent/vars"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

// Setup 安装指定组织的agent
func Setup(id, orgID uint) error {
	// 读取配置
	agents, err := repositories.GetBuildNodesByOrgId(orgID)

	if err != nil {
		return err
	}

	if agents == nil {
		return errors.New("没有找到agent配置")
	}

	for _, agent := range agents {
		if agent.ID == id {

			deploy := &kube.AppDeployConfig{
				Namespace: agent.K8sWorkerSpace,
				Name:      vars.AgentNodeName,
				Ports: []kube.Port{
					{
						ContainerPort: 80,
						ServicePort:   80,
						NodePort:      agents.NodePort,
					},
				},
				Replicas: 1,
				Image:    *conf.GetAgentImage(),
			}

			client, err := kube.NewClient(&agents.KubeConfig)
			if err != nil {
				return err
			}

			return client.Launch(deploy)
		}
	}

	return errors.New("no agent found")
}

func setupK8sNode(agent repositories.BuilderNode) {

}
