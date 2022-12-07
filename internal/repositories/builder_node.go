package repositories

import (
	"go-to-cloud/conf"
	"gorm.io/datatypes"
	"strconv"
)

type BuilderNode struct {
	Model
	BelongsTo              datatypes.JSON `json:"belongs_to" gorm:"column:belongs_to"`                             // 所属机构
	Name                   string         `json:"name" gorm:"column:name"`                                         // 节点名称
	NodeType               int            `json:"node_type" gorm:"column:node_type"`                               // 节点类型；0：k8s；2：docker；3: windows；4：linux；5：macos
	MaxWorkers             int            `json:"max_workers" gorm:"column:max_workers"`                           // 同时执行任务数量；0:不限；其他值：同时构建任务上限
	K8sWorkerSpace         string         `json:"k8s_worker_space" gorm:"column:k8s_worker_space"`                 // k8s名字空间
	K8sKubeconfigEncrypted string         `json:"k8s_kubeconfig_encrypted" gorm:"column:k8s_kubeconfig_encrypted"` // 已加密kubeconfig
}

func (m *BuilderNode) TableName() string {
	return "builder_nodes"
}

func GetBuildNodesByOrgId(orgId uint) ([]BuilderNode, error) {
	db := conf.GetDbClient()

	tx := db.Model(&BuilderNode{})

	var agents []BuilderNode
	tx = tx.Where("JSON_CONTAINS(belongs_to, ?)", strconv.Itoa(int(orgId)))
	tx = tx.Find(&agents)

	return returnWithError(agents, tx.Error)
}

func (m *BuilderNode) DecryptKubeConfig() *string {
	// TODO:
	return nil
}
