package repositories

import (
	"encoding/json"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/utils"
	"gorm.io/datatypes"
	"strconv"
	"time"
)

type BuilderNode struct {
	Model
	BelongsTo              datatypes.JSON `json:"belongs_to" gorm:"column:belongs_to"`                             // 所属机构
	Name                   string         `json:"name" gorm:"column:name"`                                         // 节点名称
	NodeType               int            `json:"node_type" gorm:"column:node_type"`                               // 节点类型；0：k8s；2：docker；3: windows；4：linux；5：macos
	MaxWorkers             int            `json:"max_workers" gorm:"column:max_workers"`                           // 同时执行任务数量；0:不限；其他值：同时构建任务上限
	K8sWorkerSpace         string         `json:"k8s_worker_space" gorm:"column:k8s_worker_space"`                 // k8s名字空间
	K8sKubeConfigEncrypted string         `json:"k8s_kubeconfig_encrypted" gorm:"column:k8s_kubeconfig_encrypted"` // 已加密kubeconfig
	k8sKubeConfigDecrypted string         `gorm:"-"`
}

func (m *BuilderNode) TableName() string {
	return "builder_nodes"
}

func GetBuildNodesById(id uint) (*BuilderNode, error) {
	db := conf.GetDbClient()

	tx := db.Model(&BuilderNode{})

	var agent BuilderNode
	tx = tx.Where("id = ?", id)
	tx = tx.First(&agent)

	return returnWithError(&agent, tx.Error)
}

func GetBuildNodesByOrgId(orgId uint) ([]BuilderNode, error) {
	db := conf.GetDbClient()

	tx := db.Model(&BuilderNode{})

	var agents []BuilderNode
	tx = tx.Where("JSON_CONTAINS(belongs_to, ?)", strconv.Itoa(int(orgId)))
	tx = tx.Find(&agents)

	return returnWithError(agents, tx.Error)
}

func (m *BuilderNode) EncryptKubeConfig() {
	m.K8sKubeConfigEncrypted = utils.Base64AesEny([]byte(m.k8sKubeConfigDecrypted))
}
func (m *BuilderNode) DecryptKubeConfig() *string {
	m.k8sKubeConfigDecrypted = utils.Base64AesEnyDecode(m.K8sKubeConfigEncrypted)
	return &m.k8sKubeConfigDecrypted
}

func buildRepoModel(model *builder.OnK8sModel, _ uint, orgs []uint, gormModel *Model) (*BuilderNode, error) {
	belongs, err := json.Marshal(orgs)
	if err != nil {
		return nil, err
	}
	repo := BuilderNode{
		Model:                  *gormModel,
		BelongsTo:              datatypes.JSON(belongs),
		Name:                   model.Name,
		NodeType:               0,
		MaxWorkers:             model.MaxWorker,
		K8sWorkerSpace:         model.Workspace,
		k8sKubeConfigDecrypted: model.KubeConfig,
	}
	repo.EncryptKubeConfig()

	return &repo, nil
}

func NewBuilderNode(node *builder.OnK8sModel, userId uint, orgs []uint) (uint, error) {
	g := &Model{
		CreatedAt: time.Now(),
	}
	repo, err := buildRepoModel(node, userId, orgs, g)
	if err != nil {
		return 0, err
	}

	tx := conf.GetDbClient()

	err = tx.Omit("updated_at").Create(&repo).Error
	if err != nil {
		return 0, err
	} else {
		return repo.ID, nil
	}
}
