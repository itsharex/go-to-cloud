package repositories

type AuthorityResource struct {
	Model
	AuthorityID int64  `json:"authority_id" gorm:"column:authority_id"`   // 权限id
	Obj         string `json:"obj" gorm:"column:obj;type: varchar(1000)"` // 资源
	Act         string `json:"act" gorm:"column:act;type: varchar(64)"`   // 方法
}

func (m *AuthorityResource) TableName() string {
	return "authority_resources"
}
