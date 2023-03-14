package users

import (
	"go-to-cloud/internal/models/user"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func GetUserList() ([]user.User, error) {
	if users, err := repositories.GetAllUser(); err != nil {
		return nil, err
	} else {
		rlt := make([]user.User, len(users))
		for i, u := range users {
			rlt[i] = user.User{
				Id:         u.ID,
				CreatedAt:  utils.JsonTime(u.CreatedAt),
				RealName:   u.RealName,
				Account:    u.Account,
				Pinyin:     u.Pinyin,
				PinyinInit: u.PinyinInit,
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
				CreatedAt: utils.JsonTime(u.CreatedAt),
				RealName:  u.RealName,
				Account:   u.Account,
			}
		}
		return rlt, nil
	}
}

func GetUserBelongs(userId uint) ([]user.Org, error) {
	if us, err := repositories.GetOrgsByUser(userId); err != nil {
		return nil, err
	} else {
		rlt := make([]user.Org, len(us))
		for i, o := range us {
			rlt[i] = user.Org{
				Id:        o.ID,
				CreatedAt: utils.JsonTime(o.CreatedAt),
				Name:      o.Name,
			}
		}
		return rlt, nil
	}
}
