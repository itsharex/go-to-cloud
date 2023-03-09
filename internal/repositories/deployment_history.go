package repositories

// DeploymentHistory k8s环境部署历史
type DeploymentHistory struct {
	Deployment
}

func (m *DeploymentHistory) TableName() string {
	return "deployments_history"
}
