package repositories

import (
	"go-to-cloud/conf"
)

type Project struct {
	Model
	Name   string `json:"name" gorm:"column:name"`
	Remark string `json:"remark" gorm:"column:remark"`
}

func (m *Project) TableName() string {
	return "projects"
}

func QueryProjectsByOrg(orgs []uint) ([]Project, error) {
	db := conf.GetDbClient()

	var projects []Project

	tx := db.Model(&Project{})

	tx = tx.Where("org.ID IN ? AND org.deleted_at IS NULL", orgs)
	err := tx.Scan(&projects).Error

	return projects, err
}
