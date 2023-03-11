package repositories

import (
	"go-to-cloud/conf"
	"gorm.io/gorm/clause"
)

// DeploymentHistory k8s环境部署历史
type DeploymentHistory struct {
	Deployment
	DeploymentId uint `json:"deploymentId" gorm:"column:deployment_id;type:bigint unsigned"`
}

func (m *DeploymentHistory) TableName() string {
	return "deployments_history"
}

func QueryDeploymentHistory(projectId, deploymentId uint) ([]DeploymentHistory, error) {
	db := conf.GetDbClient()

	var history []DeploymentHistory

	tx := db.Model(&DeploymentHistory{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("project_id = ? AND deployment_id = ?", projectId, deploymentId).Order("last_deploy_at DESC")
	err := tx.Find(&history).Error

	return returnWithError(history, err)
}
