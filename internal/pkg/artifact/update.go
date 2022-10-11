package artifact

import (
	"go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/repositories"
)

// Update 更新制品仓库
func Update(model *artifact.Artifact, userId uint, orgs map[uint]string) error {
	orgId := make([]uint, 0)
	for i := range orgs {
		orgId = append(orgId, i)
	}

	return repositories.UpdateArtifactRepo(model, userId, intersect(model.Orgs, orgId))
}
