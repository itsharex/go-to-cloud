package repositories

import (
	"encoding/json"
	"go-to-cloud/conf"
	project2 "go-to-cloud/internal/models/project"
	"gorm.io/datatypes"
	"time"
)

type Project struct {
	Model
	CreatedBy uint           `json:"createdBy" gorm:"column:created_by"`  // 仓库创建人
	BelongsTo datatypes.JSON `json:"belongsTo" gorm:"column:belongs_to;"` // 所属组织
	Name      string         `json:"name" gorm:"column:name"`
	Remark    string         `json:"remark" gorm:"column:remark"`
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
func buildProject(model *project2.DataModel, userId uint, orgs []uint, gormModel *Model) (*Project, error) {
	belongs, err := json.Marshal(orgs)
	if err != nil {
		return nil, err
	}
	return &Project{
		Model:     *gormModel,
		CreatedBy: userId,
		BelongsTo: datatypes.JSON(belongs),
		Name:      model.Name,
		Remark:    model.Remark,
	}, nil
}

func CreateProject(userId uint, orgs []uint, model project2.DataModel) (uint, error) {
	g := &Model{
		CreatedAt: time.Now(),
	}
	repo, err := buildProject(&model, userId, orgs, g)
	if err != nil {
		return 0, err
	}

	tx := conf.GetDbClient()

	// TODO: os.Env
	tx = tx.Debug()

	err = tx.Omit("updated_at").Create(&repo).Error
	return 0, err

}
