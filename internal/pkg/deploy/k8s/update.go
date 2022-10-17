package k8s

import (
	k8sModel "go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/pkg/utils"
	"go-to-cloud/internal/repositories"
)

// Update 更新代码仓库
func Update(model *k8sModel.K8s, userId uint, orgs map[uint]string) error {
	orgId := make([]uint, 0)
	for i := range orgs {
		orgId = append(orgId, i)
	}

	return repositories.UpdateK8sRepo(model, userId, utils.Intersect(model.Orgs, orgId))
}