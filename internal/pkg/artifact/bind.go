package artifact

import (
	"go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/pkg/utils"
	"go-to-cloud/internal/repositories"
)

// Bind 绑定制品仓库
func Bind(model *artifact.Artifact, userId uint, orgs map[uint]string) error {
	orgId := make([]uint, 0)
	for i := range orgs {
		orgId = append(orgId, i)
	}
	return repositories.BindArtifactRepo(model, userId, utils.Intersect(model.Orgs, orgId))
}
