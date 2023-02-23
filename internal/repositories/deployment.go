package repositories

import (
	"go-to-cloud/conf"
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"
)

// Deployment k8s环境部署方案
type Deployment struct {
	Model
	ProjectId               uint                 `json:"projectId" gorm:"project_id;type:bigint unsigned"`
	K8sNamespace            string               `json:"k8sNamespace" gorm:"column:k8s_namespace;type:varchar(20)"`
	K8sRepoId               uint                 `json:"k8sRepoId" gorm:"column:k8s_repo_id"`
	K8sRepo                 K8sRepo              `json:"-" gorm:"foreignKey:k8s_repo_id"`
	ArtifactDockerImageId   uint                 `json:"artifactDockerImageId" gorm:"column:artifact_docker_image_id;type:bigint unsigned"`
	ArtifactDockerImageRepo ArtifactDockerImages `json:"-" gorm:"foreignKey:artifact_docker_image_id"`
	Ports                   datatypes.JSON       `json:"ports" gorm:"column:ports"` // 端口{containerPort: 80, servicePort: 80, nodePort: 30080, portName: 'http'}
	Cpus                    uint                 `json:"cpus" gorm:"column:cpus;type:int unsigned"`
	Env                     datatypes.JSON       `json:"env" gorm:"column:env;type:text"`                                             // 环境变量，json形式
	Replicas                uint                 `json:"replicas" gorm:"column:replicas;type:int unsigned"`                           // 副本数量
	Liveness                string               `json:"liveness" gorm:"column:liveness;type: varchar(500)"`                          // 存活检查地址
	Readiness               string               `json:"readiness" gorm:"column:readiness;type: varchar(500)"`                        // 就绪检查地址
	RollingMaxSurge         uint                 `json:"rollingMaxSurge" gorm:"column:rolling_max_surge;type:int unsigned"`           // 滚动发布策略：激增数量上限（1～100）
	RollingMaxUnavailable   uint                 `json:"rollingMaxUnavailable" gorm:"rolling_max_unavailable;type:int unsigned"`      // 滚动发布策略：最大不可用上限(1~100)
	ResourceLimitCpuRequest uint                 `json:"resourceLimitCpuRequest" gorm:"resource_limit_cpu_request;type:int unsigned"` // 资源限制：cpu分配数量，单位m
	ResourceLimitCpuLimits  uint                 `json:"resourceLimitCpuLimits" gorm:"resource_limit_cpu_limits;type:int unsigned"`   // 资源限制：cpu分配上限，单位m
	ResourceLimitMemRequest uint                 `json:"resourceLimitMemRequest" gorm:"resource_limit_mem_request;type:int unsigned"` // 资源限制：内在分配数量，单位Mi
	ResourceLimitMemLimits  uint                 `json:"resourceLimitMemLimits" gorm:"resource_limit_mem_limits;type:int unsigned"`   // 资源限制：内在分配上限，单位Mi
	NodeSelector            datatypes.JSON       `json:"nodeSelector" gorm:"node_selector;"`                                          // 节点选择，json，[{"labelName": "labelValue"}]
}

func (m *Deployment) TableName() string {
	return "deployments"
}

func QueryDeploymentsByProjectId(projectId uint) ([]Deployment, error) {
	db := conf.GetDbClient()

	var deployments []Deployment

	tx := db.Model(&Deployment{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("project_id = ?", projectId)
	err := tx.Find(&deployments).Error

	return returnWithError(deployments, err)
}

func CreateDeployment(deployment *Deployment) error {
	db := conf.GetDbClient()

	tx := db.Model(&Deployment{})

	return tx.Create(deployment).Error
}
