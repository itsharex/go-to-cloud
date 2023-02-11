package repositories

import (
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/utils"
)

type PipelineHistory struct {
	Model
	PipelineID    uint                    `json:"pipeline_id" gorm:"column:pipeline_id;index:pipeline_history_pipeline_id_index"`
	ProjectID     uint                    `json:"project_id" gorm:"column:project_id;index:pipeline_history_project_id_index"`
	Name          string                  `json:"name" gorm:"column:name;type:nvarchar(64)"` // 计划名称
	Env           string                  `json:"env" gorm:"column:env;type:nvarchar(64)"`   // 运行环境(模板), e.g. dotnet:6; go:1.17
	SourceCodeID  uint                    `json:"source_code_id" gorm:"column:source_code_id"`
	Branch        string                  `json:"branch" gorm:"column:branch;type:nvarchar(64)"` // 分支名称
	Params        string                  `json:"params" gorm:"column:params"`                   // 本次运行的参数(json格式）
	CreatedBy     uint                    `json:"created_by" gorm:"column:created_by"`           // 构建人
	Remark        string                  `json:"remark" gorm:"column:remark"`
	BuildAt       *utils.JsonTime         `json:"build_at" gorm:"column:build_at"`               // 最近一次运行时间
	LastRunResult pipeline.BuildingResult `json:"last_run_result" gorm:"column:last_run_result"` // 最近一次运行结果; 1：成功；2：取消；3：失败；0：从未执行
}

func (m *PipelineHistory) TableName() string {
	return "pipeline_history"
}
