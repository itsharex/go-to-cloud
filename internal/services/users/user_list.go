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
				Id:        u.ID,
				CreatedAt: utils.JsonTime(u.CreatedAt),
				Name:      u.RealName,
				Account:   u.Account,
				Shortcut:  u.Shortcut,
			}
		}
		return rlt, nil
	}
}
