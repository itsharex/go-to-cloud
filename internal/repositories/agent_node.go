package repositories

import (
	"go-to-cloud/conf"
	"gorm.io/datatypes"
	"strconv"
)

type AgentNode struct {
	Model
	BelongsTo  datatypes.JSON `json:"belongsTo" gorm:"column:belongs_to;"` // 所属组织
	KubeConfig string         `json:"kube_config" gorm:"column:kube_config"`
	Namespace  string         `json:"namespace" gorm:"column:namespace"`
	NodePort   int            `json:"node_port" gorm:"column:node_port"` // agent服务端口
	CreatedBy  uint64         `json:"created_by" gorm:"column:created_by"`
	Remark     string         `json:"remark" gorm:"column:remark"`
}

func (m *AgentNode) TableName() string {
	return "agent_node"
}

func GetAgentByOrgId(orgId uint) (*AgentNode, error) {
	db := conf.GetDbClient()

	tx := db.Model(&AgentNode{})
	if conf.Environment.IsDevelopment() {
		tx = tx.Debug()
	}

	var agent AgentNode
	tx = tx.Where("JSON_CONTAINS(belongs_to, ?)", strconv.Itoa(int(orgId)))
	tx = tx.First(&agent)

	return returnWithError(&agent, tx.Error)
}
