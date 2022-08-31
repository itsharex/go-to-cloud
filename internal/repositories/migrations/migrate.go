package migrations

import (
	"go-to-cloud/conf"
	"gorm.io/gorm"
)

type Migration interface {
	Up(db *gorm.DB)
	Down(db *gorm.DB)
}

// Migrate 数据库变更同步
func Migrate() {
	db := conf.GetDbClient()

	var up Migration

	up = &Migration20220831{}
	up.Up(db)
}

func Rollback() {
	db := conf.GetDbClient()

	var up Migration

	up = &Migration20220831{}
	up.Down(db)
}
