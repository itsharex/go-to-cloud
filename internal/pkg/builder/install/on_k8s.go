package install

import (
	"go-to-cloud/internal/agent"
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func OnK8s(model *builder.OnK8sModel, userId uint, orgId []uint) error {
	id, err := repositories.NewBuilderNode(model, userId, utils.Intersect(model.Orgs, orgId))
	if err != nil {
		return err
	}

	// 部署到k8s
	return agent.Setup(id)
}
