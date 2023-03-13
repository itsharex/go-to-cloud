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
				Key:        u.ID,
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
