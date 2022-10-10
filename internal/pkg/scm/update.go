package scm

import (
	"go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/repositories"
)

// Update 更新代码仓库
func Update(model *scm.Scm, userId uint, orgs map[uint]string) error {
	orgId := make([]uint, 0)
	for i := range orgs {
		orgId = append(orgId, i)
	}

	return repositories.UpdateCodeRepo(model, userId, intersect(model.Orgs, orgId))
}
