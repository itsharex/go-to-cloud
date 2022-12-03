package repositories

import (
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/build"
	"time"
)

type ArtifactScript struct {
	Dockerfile string `json:"dockerfile"`
	Registry   string `json:"registry"`
	IsSecurity bool   `json:"isSecurity"`
	Account    string `json:"account"`
	Password   string `json:"password"`
}

type CiPlan struct {
	Model
	CiPlanSteps     []CiPlanSteps `json:"-" gorm:"foreignKey:ci_plan_id"`
	ProjectID       uint          `json:"project_id" gorm:"column:project_id"`
	Name            string        `json:"name" gorm:"column:name"` // 计划名称
	Env             string        `json:"env" gorm:"column:env"`   // 构建环境, e.g. dotnet:6; go:1.17
	SourceCodeID    uint          `json:"source_code_id" gorm:"column:source_code_id"`
	Branch          string        `json:"branch" gorm:"column:branch"` // 分支名称
	CreatedBy       uint          `json:"created_by" gorm:"column:created_by"`
	Remark          string        `json:"remark" gorm:"column:remark"`
	LastBuildAt     time.Time     `json:"last_build_at" gorm:"column:last_build_at"`         // 最近一次构建时间
	LastBuildResult int           `json:"last_build_result" gorm:"column:last_build_result"` // 最近一次构建结果; 1：成功；2：取消；3：失败；0：从未执行
}

func (m *CiPlan) TableName() string {
	return "ci_plan"
}

// NewPlan 新建构建计划
func NewPlan(projectId uint, currentUserId uint, model *build.PlanModel) (err error) {
	steps := make(steps, 0)
	sort := 0
	err = steps.qaStep(model, &sort)
	if err != nil {
		return err
	}
	err = steps.artifactStep(model, &sort)
	if err != nil {
		return err
	}

	plan := CiPlan{
		ProjectID:       projectId,
		Name:            model.Name,
		Env:             model.BuildEnv,
		SourceCodeID:    model.SourceCodeId,
		Branch:          model.Branch,
		CreatedBy:       currentUserId,
		Remark:          model.Remark,
		LastBuildResult: 0,
	}

	plan.CiPlanSteps = steps

	tx := conf.GetDbClient()

	err = tx.Create(plan).Error

	return
}
