package repositories

import (
	"errors"
	"go-to-cloud/conf"
)

type ProjectSourceCode struct {
	Model
	Project   Project `json:"-" gorm:"foreignKey:project_id"`
	ProjectID uint    `json:"project_id" gorm:"column:project_id"` // 所属项目
	GitUrl    string  `json:"git_url" gorm:"column:git_url"`       // git地址
	CreatedBy uint    `json:"created_by" gorm:"column:created_by"`
}

func (m *ProjectSourceCode) TableName() string {
	return "project_source_code"
}

func UpsertProjectSourceCode(projectId, userId uint, url *string) error {
	db := conf.GetDbClient()

	tx := db.Model(&ProjectSourceCode{})

	if conf.Environment.IsDevelopment() {
		tx = tx.Debug()
	}

	tx = tx.FirstOrCreate(&ProjectSourceCode{
		ProjectID: projectId,
		GitUrl:    *url,
	})

	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("already exists")
	}

	return nil
}
