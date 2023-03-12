package repositories

import "go-to-cloud/conf"

// Org 组织
// 组织拥有基础设施和用户
type Org struct {
	Model
	Name   string  `json:"name" gorm:"column:name;not null;"` // 组织名称
	Users  []*User `gorm:"many2many:orgs_users_rel;"`
	Remark string  `json:"remark" gorm:"column:remark;type:nvarchar(1024);"`
}

func (m *Org) TableName() string {
	return "org"
}

func GetOrgs() ([]Org, error) {
	db := conf.GetDbClient()

	var org []Org
	err := db.Preload("Users").Find(&org).Error

	return org, err
}
