package k8s

import (
	"go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/pkg/utils"
	"go-to-cloud/internal/repositories"
)

// Bind 绑定代码仓库
func Bind(model *k8s.K8s, userId uint, orgs map[uint]string) error {
	orgId := make([]uint, 0)
	for i := range orgs {
		orgId = append(orgId, i)
	}
	return repositories.BindK8sRepo(model, userId, utils.Intersect(model.Orgs, orgId))
}
