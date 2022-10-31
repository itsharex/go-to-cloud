package repositories

type GitRepo struct {
	Model
	Project   Project
	ProjectId uint64 `json:"project_id" gorm:"column：project_id"`
	Url       string `json:"url" gorm:"column:url"` // SCM平台地址（非项目仓库地址）
}

func (m *GitRepo) TableName() string {
	return "git_repo"
}
