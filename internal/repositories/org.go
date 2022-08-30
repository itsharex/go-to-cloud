package repositories

// Org 组织
type Org struct {
	AddOn
	Name            string           `json:"name" gorm:"column:name"` // 组织名称
	Infrastructures []Infrastructure `gorm:"foreignKey:ID"`           // 基础设施
	Users           []*User          `gorm:"many2many:orgs_users_rel;"`
}

func (m *Org) TableName() string {
	return "org"
}
