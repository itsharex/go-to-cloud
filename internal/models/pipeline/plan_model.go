package pipeline

import (
	"errors"
	"go-to-cloud/internal/utils"
	"strings"
)

type PlanStepType int

// 1:运行单测；2：运行lint；3：生成文档；4：生成镜像；5：部署；0：其他cli命令
const (
	Cli       PlanStepType = 0
	UnitTest               = 1
	LintCheck              = 2
	Doc                    = 3
	Image                  = 4
	Deploy                 = 5
)

type BuildingResult int

const (
	NeverBuild        BuildingResult = 0
	BuildingSuccess   BuildingResult = 1
	BuildingInterrupt BuildingResult = 2
	BuildingFailed    BuildingResult = 3
)

// PlanModel 构建计划模型
type PlanModel struct {
	Id              uint    `json:"id"`
	Name            string  `json:"name"`
	Env             string  `json:"buildEnv"`
	SourceCodeID    uint    `json:"source_code_id"`
	Branch          string  `json:"branch"`
	QaEnabled       bool    `json:"qa_enabled"`
	UnitTest        *string `json:"unit_test"`
	LintCheck       *string `json:"lint_check"`
	ArtifactEnabled bool    `json:"artifact_enabled"`
	Dockerfile      *string `json:"dockerfile"`
	ArtifactRepoId  *uint   `json:"artifact_repo_id"`
	DeployEnabled   bool    `json:"deploy_enabled"`
	Remark          string  `json:"remark"`
}

type PlanCardModel struct {
	PlanModel
	LastBuildAt     *utils.JsonTime `json:"lastBuildAt"`
	LastBuildResult BuildingResult  `json:"lastBuildResult"`
}

func (m *PlanModel) Valid() error {
	if len(strings.TrimSpace(m.Name)) == 0 {
		return errors.New("name is empty")
	}
	if len(strings.TrimSpace(m.Env)) == 0 {
		return errors.New("build env is not selected")
	}
	if len(strings.TrimSpace(m.Branch)) == 0 {
		return errors.New("branch is not selected")
	}
	if m.SourceCodeID == 0 {
		return errors.New("source code is not selected")
	}
	return nil
}
