package install

import (
	"errors"
	"go-to-cloud/internal/agent"
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
	"strings"
)

func OnK8s(model *builder.OnK8sModel, userId uint, orgId []uint) error {
	// 检查目标主机是否已经存在同名namespace，如果存在，则不允许创建
	client, err := kube.NewClient(&model.KubeConfig)
	if err != nil {
		return err
	}
	ns, err := client.GetAllNamespaces(true)
	if err != nil {
		return err
	}
	for _, n := range ns {
		if strings.EqualFold(n, model.Workspace) {
			return errors.New("agent exists, change workspace and try again")
		}
	}

	id, err := repositories.NewBuilderNode(model, userId, utils.Intersect(model.Orgs, orgId))
	if err != nil {
		return err
	}

	// 部署到k8s
	return agent.Setup(id)
}
