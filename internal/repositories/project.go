package repositories

import (
	"errors"
	"fmt"
	"go-to-cloud/conf"
	project2 "go-to-cloud/internal/models/project"
	"time"
)

type Project struct {
	Model
	CreatedBy uint   `json:"createdBy" gorm:"column:created_by"` // 仓库创建人
	OrgId     uint   `json:"orgId" gorm:"column:org_id;"`        // 所属组织
	Name      string `json:"name" gorm:"column:name"`
	Remark    string `json:"remark" gorm:"column:remark"`
}

func (m *Project) TableName() string {
	return "projects"
}

func QueryProjectsByOrg(orgs []uint) ([]Project, error) {
	db := conf.GetDbClient()

	var projects []Project

	tx := db.Model(&Project{})

	if conf.Environment.IsDevelopment() {
		tx = tx.Debug()
	}

	tx = tx.Select("projects.*, org.Id AS orgId, org.Name AS orgName")
	tx = tx.Joins("INNER JOIN org ON projects.org_id = org.ID")
	tx = tx.Where("org.ID IN ? AND org.deleted_at IS NULL", orgs)
	err := tx.Scan(&projects).Error

	return projects, err
}
func buildProject(model *project2.DataModel, userId uint, orgId uint, gormModel *Model) (*Project, error) {
	return &Project{
		Model:     *gormModel,
		CreatedBy: userId,
		OrgId:     orgId,
		Name:      model.Name,
		Remark:    model.Remark,
	}, nil
}

func CreateProject(userId uint, orgId uint, model project2.DataModel) (uint, error) {
	g := &Model{
		CreatedAt: time.Now(),
	}

	repo, err := buildProject(&model, userId, orgId, g)
	if err != nil {
		return 0, err
	}

	tx := conf.GetDbClient()

	if conf.Environment.IsDevelopment() {
		tx = tx.Debug()
	}

	var total int64
	err = tx.Model(&Project{}).Where("name = ? AND org_id = ?", model.Name, orgId).Count(&total).Error
	if err != nil {
		return 0, err
	}
	if total > 0 {
		return 0, errors.New(fmt.Sprintf("project '%s' already exists", model.Name))
	}

	err = tx.Omit("updated_at").Create(&repo).Error
	return 0, err

}
