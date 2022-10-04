package scm

import (
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/repositories"
)

// Bind 绑定代码仓库
func Bind(model *models.Scm, userId int64, orgs map[int64]string) error {
	orgId := make([]int64, 0)
	for i, _ := range orgs {
		orgId = append(orgId, i)
	}
	return repositories.BindCodeRepo(model, userId, orgId)
}
