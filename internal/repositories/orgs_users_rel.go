package repositories

import "go-to-cloud/conf"

// FetchUsersByOrg 获取指定组织下的用户
// orgId：所属组织
func FetchUsersByOrg(orgId uint) ([]User, error) {
	db := conf.GetDbClient()

	var users []User

	err := db.Debug().Preload("orgs", orgId).Find(&users).Error

	return users, err
}
