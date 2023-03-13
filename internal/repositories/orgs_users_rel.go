package repositories

import (
	"go-to-cloud/conf"
)

// GetUsersByOrg 获取指定组织下的用户
// orgId：所属组织
func GetUsersByOrg(orgId uint) ([]*User, error) {
	db := conf.GetDbClient()

	var org Org
	err := db.Preload("Users").Where([]uint{orgId}).First(&org).Error

	return org.Users, err
}

func AddMembersToOrg(orgId uint, memebers []uint) error {
	db := conf.GetDbClient()

	var org Org
	if err := db.Preload("Users").Where([]uint{orgId}).First(&org).Error; err != nil {
		return err
	} else {

	}
}
