package build

// PlanModel 构建计划模型
type PlanModel struct {
	Name            string `json:"name"`
	BuildEnv        string `json:"buildEnv"`
	SourceCodeId    uint   `json:"source_code_id"`
	Branch          string `json:"branch"`
	QaEnabled       bool   `json:"qa_enabled"`
	UnitTest        string `json:"unit_test"`
	LintCheck       string `json:"lint_check"`
	ArtifactEnabled bool   `json:"artifact_enabled"`
	Dockerfile      string `json:"dockerfile"`
	ArtifactRepoId  uint   `json:"artifact_repo_id"`
	DeployEnabled   bool   `json:"deploy_enabled"`
	Remark          string `json:"remark"`
}

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
