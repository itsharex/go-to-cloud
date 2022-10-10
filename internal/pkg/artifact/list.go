package artifact

import (
	"go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/repositories"
)

// List 读取制品仓库
// @Params:
//
//	orgs: 当前用户所在组织
//	query: 查询条件
func List(orgs []uint, query *artifact.Query) ([]artifact.Artifact, error) {
	var orgId []uint
	if len(query.Orgs) == 0 {
		//	默认取当前用户所属全体组织
		orgId = orgs
	} else {
		// 	计算查询条件中的所属组织与当前用户所属组织的交集
		orgId = intersect(orgs, query.Orgs)
	}

	if merged, err := repositories.QueryArtifactRepo(orgId, query.Name); err != nil {
		return nil, err
	} else {
		rlt := make([]artifact.Artifact, len(merged))
		for i, m := range merged {
			// TODO: uncomment
			//orgLites := make([]artifact.OrgLite, len(m.Org))
			//for i, lite := range m.Org {
			//	orgLites[i] = artifact.OrgLite{
			//		OrgId:   lite.OrgId,
			//		OrgName: lite.OrgName,
			//	}
			//}
			rlt[i] = artifact.Artifact{
				Testing: artifact.Testing{
					Id:         m.ID,
					Type:       artifact.Type(m.ArtifactOrigin),
					IsSecurity: m.IsSecurity != 0,
					Url:        m.Url,
					User:       m.Account,
					Password:   m.Password,
				},
				Name: m.Name,
				// TODO: uncomment
				//OrgLites:  orgLites,
				Remark:    m.Remark,
				UpdatedAt: m.UpdatedAt.Format("2006-01-02"),
			}
		}
		return rlt, err
	}
}
