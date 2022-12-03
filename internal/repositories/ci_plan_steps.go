package repositories

import (
	"encoding/json"
	"errors"
	"go-to-cloud/internal/models/build"
	"strings"
)

type CiPlanSteps struct {
	Model
	CiPlan   CiPlan `json:"-" gorm:"foreignKey:ci_plan_id"`
	CiPlanID int64  `json:"ci_plan_id" gorm:"column:ci_plan_id"`
	Sort     int    `json:"sort" gorm:"column:sort"`               // 执行顺序
	Name     string `json:"name" gorm:"column:name"`               // 步骤名称
	Script   string `json:"script" gorm:"column:script;type:text"` // 步骤脚本；当步骤类型为(5)部署时，script表示deployment和service的yml；为(4)生成制品时；由制品类型决定内容
	Type     int    `json:"type" gorm:"column:type"`               // 节点类型; 1:运行单测；2：运行lint；3：生成文档；4：生成镜像；5：部署；0：其他cli命令
}

func (m *CiPlanSteps) TableName() string {
	return "ci_plan_steps"
}

type steps []CiPlanSteps

func (steps *steps) qaStep(model *build.PlanModel, sort *int) error {
	if model.QaEnabled {
		if len(strings.TrimSpace(model.UnitTest)) > 0 {
			*steps = append(*steps, CiPlanSteps{
				Sort:   *sort,
				Name:   "单元测试",
				Script: model.UnitTest,
				Type:   build.UnitTest,
			})
			*sort++
		}
		if len(strings.TrimSpace(model.LintCheck)) > 0 {
			*steps = append(*steps, CiPlanSteps{
				Sort:   *sort,
				Name:   "Lint检查",
				Script: model.LintCheck,
				Type:   build.LintCheck,
			})
			*sort++
		}
	}
	return nil
}

func (steps *steps) artifactStep(model *build.PlanModel, sort *int) error {
	if model.ArtifactEnabled {
		if url, account, password, isSecurity, origin, err := GetArtifactRepoByID(model.ArtifactRepoId); err != nil {
			return err
		} else {
			if origin != 1 {
				return errors.New("not docker registry")
			}
			script, _ := json.Marshal(ArtifactScript{
				Dockerfile: model.Dockerfile,
				Registry:   *url,
				IsSecurity: isSecurity,
				Account:    *account,
				Password:   *password,
			})
			*steps = append(*steps, CiPlanSteps{
				Sort:   *sort,
				Name:   "镜像制品",
				Script: string(script),
				Type:   build.Image,
			})
			*sort++
		}
	}
	return nil
}
