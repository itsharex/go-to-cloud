package repositories

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
	"time"
)

// User 登录账户
type User struct {
	gorm.Model
	Account        string     `json:"account" gorm:"column:account;not null;"`   // 账号
	HashedPassword string     `json:"-" gorm:"column:password;not null;"`        // 登录密码
	Email          string     `json:"email" gorm:"column:email"`                 // 邮箱
	Mobile         string     `json:"mobile" gorm:"column:mobile"`               // 联系电话
	LastLoginAt    *time.Time `json:"last_login_at" gorm:"column:last_login_at"` // 上次登录时间
	Orgs           []*Org     `gorm:"many2many:orgs_users_rel"`
}

func (m *User) TableName() string {
	return "users"
}

// SetPassword 加密密码
func (m *User) SetPassword(origPassword *string) error {
	if len(strings.Trim(*origPassword, " ")) == 0 {
		return errors.New("密码不允许为空")
	}
	lowerPassword := strings.ToLower(*origPassword)
	if hashBytes, err := bcrypt.GenerateFromPassword([]byte(lowerPassword), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		m.HashedPassword = string(hashBytes)
		return nil
	}
}

// ComparePassword 比较密码
func (m *User) ComparePassword(password *string) bool {
	lowerPassword := strings.ToLower(*password)
	return nil == bcrypt.CompareHashAndPassword([]byte(m.HashedPassword), []byte(lowerPassword))
}
