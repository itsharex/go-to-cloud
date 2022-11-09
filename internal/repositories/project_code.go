package repositories

import (
	"errors"
	"go-to-cloud/conf"
	"gorm.io/gorm/clause"
)

type ProjectSourceCode struct {
	Model
	Project     Project  `json:"-" gorm:"foreignKey:project_id"`
	ProjectID   uint     `json:"project_id" gorm:"column:project_id"` // 所属项目
	CodeRepo    CodeRepo `json:"-" gorm:"foreignKey:code_repo_id"`
	CodeRepoID  uint     `json:"code_repo_id" gorm:"column:code_repo_id"`
	GitUrl      string   `json:"git_url" gorm:"column:git_url"` // git地址
	CreatedUser User     `json:"-" gorm:"foreignKey:created_by"`
	CreatedBy   uint     `json:"created_by" gorm:"column:created_by"`
}

func (m *ProjectSourceCode) TableName() string {
	return "project_source_code"
}

func UpsertProjectSourceCode(projectId, codeRepoId, userId uint, url *string) error {
	db := conf.GetDbClient()

	tx := db.Model(&ProjectSourceCode{})

	if conf.Environment.IsDevelopment() {
		tx = tx.Debug()
	}

	tx = tx.FirstOrCreate(&ProjectSourceCode{
		CodeRepoID: codeRepoId,
		ProjectID:  projectId,
		GitUrl:     *url,
		CreatedBy:  userId,
	})

	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("already exists")
	}

	return nil
}

func QueryProjectSourceCode(projectId uint) ([]ProjectSourceCode, error) {
	db := conf.GetDbClient()
	tx := db.Model(&ProjectSourceCode{})
	tx = tx.Preload(clause.Associations)

	if conf.Environment.IsDevelopment() {
		tx = tx.Debug()
	}

	var rlt []ProjectSourceCode
	tx = tx.Where("project_id = ?", projectId).Find(&rlt)

	return rlt, tx.Error
}
