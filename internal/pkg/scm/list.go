package scm

import (
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/repositories"
)

// List 读取指定用户、组织下的仓库
func List(userId int64, orgs []int64, query *models.ScmQuery) ([]models.Scm, error) {
	var orgId []int64
	if len(query.Orgs) == 0 {
		//	默认取当前用户所属全体组织
		orgId = orgs
	} else {
		// 	计算查询条件中的所属组织与当前用户所属组织的交集
		orgId = intersect(orgs, query.Orgs)
	}

	if _, err := repositories.ReadCodeRepo(orgId, query.Name); err != nil {
		return nil, err
	} else {
		return nil, err
	}
}

func intersect(orgA, orgB []int64) []int64 {
	counter := make(map[int64]int)
	rlt := make([]int64, 0)
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
