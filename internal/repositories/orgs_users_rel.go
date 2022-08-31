package repositories

import "go-to-cloud/conf"

// FetchUsersByOrg 获取指定组织下的用户
// orgId：所属组织
func FetchUsersByOrg(orgId uint) ([]*User, error) {
	db := conf.GetDbClient()

	var org Org
	err := db.Preload("Users").Where([]uint{orgId}).First(&org).Error

	return org.Users, err
}
