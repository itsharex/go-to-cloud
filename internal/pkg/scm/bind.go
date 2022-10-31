package scm

import (
	"go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/pkg/utils"
	"go-to-cloud/internal/repositories"
)

// Bind 绑定代码仓库
func Bind(model *scm.Scm, userId uint, orgId []uint) error {
	return repositories.BindCodeRepo(model, userId, utils.Intersect(model.Orgs, orgId))
}
