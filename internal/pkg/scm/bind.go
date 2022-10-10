package scm

import (
	"go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/repositories"
)

// Bind 绑定代码仓库
func Bind(model *scm.Scm, userId uint, orgs map[uint]string) error {
	orgId := make([]uint, 0)
	for i := range orgs {
		orgId = append(orgId, i)
	}
	return repositories.BindCodeRepo(model, userId, intersect(model.Orgs, orgId))
}
