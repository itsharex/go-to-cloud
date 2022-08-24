package repositories

import "time"

type Infrastructures struct {
	ID            int       `json:"id" gorm:"column:id"`
	Remark        string    `json:"remark" gorm:"column:remark"`                 // 设施备注
	EncodedConfig string    `json:"encoded_config" gorm:"column:encoded_config"` // 编码后的配置内容
	Type          int       `json:"type" gorm:"column:type"`                     // 设施分类；1：k8s；2：registry；
	IsDeleted     int8      `json:"is_deleted" gorm:"column:is_deleted"`         // 删除标记；1：删除；0：未删除
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at"`         // 创建时间
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at"`         // 更新时间
}

func (m *Infrastructures) TableName() string {
	return "infrastructures"
}
