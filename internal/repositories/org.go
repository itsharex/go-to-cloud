package repositories

// Org 组织
type Org struct {
	AddOn
	Name string `json:"name" gorm:"column:name"` // 组织名称
}

func (m *Org) TableName() string {
	return "org"
}
