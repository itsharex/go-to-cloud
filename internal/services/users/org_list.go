package users

import (
	"go-to-cloud/internal/models/user"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func GetOrgList() ([]user.Org, error) {
	if orgs, err := repositories.GetOrgs(); err != nil {
		return nil, err
	} else {
		rlt := make([]user.Org, len(orgs))
		for i, org := range orgs {
			rlt[i] = user.Org{
				Id:          org.ID,
				CreatedAt:   utils.JsonTime(org.CreatedAt),
				Name:        org.Name,
				MemberCount: uint(len(org.Users)),
				Remark:      org.Remark,
			}
		}
		return rlt, nil
	}
}
