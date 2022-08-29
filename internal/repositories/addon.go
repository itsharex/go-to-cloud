package repositories

import (
	"gorm.io/gorm"
)

type AddOn struct {
	gorm.Model
	IsDeleted int8 `json:"is_deleted" gorm:"column:is_deleted"` // 删除标记；1：删除；0：未删除
}
