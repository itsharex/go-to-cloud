package agent

import (
	"errors"
	"go-to-cloud/conf"
	"go-to-cloud/internal/agent/vars"
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

// Setup 安装指定组织的agent
func Setup(id uint) error {
	// 读取配置
	agent, err := repositories.GetBuildNodesById(id)

	if err != nil {
		return err
	}

	if agent == nil {
		return errors.New("没有找到agent配置")
	}

	if agent.NodeType == int(builder.K8s) {
		return setupK8sNode(agent)
	}

	return errors.New("no agent found")
}

func setupK8sNode(agent *repositories.BuilderNode) error {
	deploy := &kube.AppDeployConfig{
		Namespace: agent.K8sWorkerSpace,
		Name:      vars.AgentNodeName,
		Replicas:  1,
		Image:     *conf.GetAgentImage(),
		Env: []kube.EnvVar{
			{
				Name:  "GoToCloud-Server",
				Value: conf.GetServerGrpcHost().Host,
			},
		},
	}

	client, err := kube.NewClient(agent.DecryptKubeConfig())
	if err != nil {
		return err
	}

	return client.Launch(deploy)
}
