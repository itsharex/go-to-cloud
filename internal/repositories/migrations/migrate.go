package migrations

import (
	"gorm.io/gorm"
)

type Migration interface {
	Up(db *gorm.DB)
	Down(db *gorm.DB)
}

var migrations []Migration

func init() {
	// 迁移对象必需按从旧到新的顺序添加
	migrations = []Migration{
		&Migration20220831{},
		&migration20220921{},
		&migration20221004{},
	}
}

// Migrate 数据库变更同步
func Migrate(db *gorm.DB) {

	for i := 0; i < len(migrations); i++ {
		migrations[i].Up(db)
	}
}

// Rollback 数据库变更回滚
func Rollback(db *gorm.DB) {

	for i := len(migrations) - 1; i >= 0; i-- {
		migrations[i].Down(db)
	}
}
