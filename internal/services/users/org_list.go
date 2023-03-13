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

func GetUsersByOrg(orgId uint) ([]user.User, error) {
	if us, err := repositories.GetUsersByOrg(orgId); err != nil {
		return nil, err
	} else {
		rlt := make([]user.User, len(us))
		for i, u := range us {
			rlt[i] = user.User{
				Id:        u.ID,
				Key:       u.ID,
				CreatedAt: utils.JsonTime(u.CreatedAt),
				RealName:  u.RealName,
				Account:   u.Account,
			}
		}
		return rlt, nil
	}
}

// JoinOrg 加入组织，如果当前成员在users中不存在，则移除
func JoinOrg(orgId uint, users []uint) error {
	if us, err := repositories.GetUsersByOrg(orgId); err != nil {
		return err
	} else {
		old := make([]uint, len(us))
		for i, u := range us {
			old[i] = u.ID
		}

		oldSet := utils.New(old...)
		newSet := utils.New(users...)

		delSet := utils.Minus(oldSet, newSet)      // 差集：移除
		comSet := utils.Complement(oldSet, newSet) // 补集：追加

		return repositories.UpdateMembersToOrg(orgId, utils.List(comSet), utils.List(delSet))
	}
}
