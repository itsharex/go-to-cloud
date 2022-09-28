package repositories

import "gorm.io/gorm"

type CodeRepo struct {
	gorm.Model
}

func (m *CodeRepo) TableName() string {
	return "coderepo"
}
