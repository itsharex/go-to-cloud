package scm

import (
	"go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/repositories"
)

// List 读取仓库
// @Params:
//
//	orgs: 当前用户所在组织
//	query: 查询条件
func List(orgs []uint, query *scm.Query) ([]scm.Scm, error) {
	var orgId []uint
	if len(query.Orgs) == 0 {
		//	默认取当前用户所属全体组织
		orgId = orgs
	} else {
		// 	计算查询条件中的所属组织与当前用户所属组织的交集
		orgId = intersect(orgs, query.Orgs)
	}

	if merged, err := repositories.QueryCodeRepo(orgId, query.Name, &query.Pager); err != nil {
		return nil, err
	} else {
		rlt := make([]scm.Scm, len(merged))
		for i, m := range merged {
			orgLites := make([]scm.OrgLite, len(m.Org))
			for i, lite := range m.Org {
				orgLites[i] = scm.OrgLite{
					OrgId:   lite.OrgId,
					OrgName: lite.OrgName,
				}
			}
			rlt[i] = scm.Scm{
				Testing: scm.Testing{
					Id:       m.ID,
					Origin:   scm.Type(m.ScmOrigin),
					IsPublic: m.IsPublic != 0,
					Url:      m.Url,
					Token:    &merged[i].AccessToken,
				},
				Name:      m.Name,
				OrgLites:  orgLites,
				Remark:    m.Remark,
				UpdatedAt: m.UpdatedAt.Format("2006-01-02"),
			}
		}
		return rlt, err
	}
}

func intersect(orgA, orgB []uint) []uint {
	counter := make(map[uint]int)
	rlt := make([]uint, 0)
	for _, a := range orgA {
		counter[a]++
	}
	for _, b := range orgB {
		sz, _ := counter[b]
		if sz == 1 {
			rlt = append(rlt, b)
		}
	}
	return rlt
}
