package project

import "time"

type SourceCodeModel struct {
	CodeRepoId uint   `json:"codeRepoId"` // 仓库ID
	Url        string `json:"url"`        // 代码地址
}

type SourceCodeImportedModel struct {
	SourceCodeModel
	CodeRepoOrigin int       `json:"codeRepoOrigin"`
	Id             uint      `json:"id"`        // 代码ID
	CreatedBy      string    `json:"createdBy"` // 导入人
	CreatedAt      time.Time `json:"createdAt"` // 导入时间
}
