package repositories

import (
	"gorm.io/gorm"
	"time"
)

// User 登录账户
type User struct {
	gorm.Model
	Account     string     `json:"account" gorm:"column:account;not null;"`   // 账号
	Password    string     `json:"password" gorm:"column:password;not null;"` // 登录密码
	Email       string     `json:"email" gorm:"column:email"`                 // 邮箱
	Mobile      string     `json:"mobile" gorm:"column:mobile"`               // 联系电话
	LastLoginAt *time.Time `json:"last_login_at" gorm:"column:last_login_at"` // 上次登录时间
	Orgs        []*Org     `gorm:"many2many:orgs_users_rel"`
}

func (m *User) TableName() string {
	return "users"
}
