package repositories

import (
	"encoding/base64"
	"go-to-cloud/conf"
	"go-to-cloud/internal/utils"
	"gorm.io/gorm"
)

// Infrastructure 基础设施
type Infrastructure struct {
	gorm.Model
	OrgId         int64      `json:"org_id" gorm:"column:org_id;not null"`                 // 所属组织
	Remark        string     `json:"remark" gorm:"column:remark"`                          // 设施备注
	EncodedConfig string     `json:"encoded_config" gorm:"column:encoded_config;not null"` // 编码后的配置内容
	Type          InfraTypes `json:"type" gorm:"column:type;not null"`                     // 设施分类；1：k8s；2：registry；
	Config        *string    `gorm:"-"`                                                    // Config原文
}

func (m *Infrastructure) TableName() string {
	return "infrastructures"
}

func (m *Infrastructure) DecodeConfig() {
	t, _ := base64.StdEncoding.DecodeString(m.EncodedConfig)
	*m.Config = string(utils.AesEny(t))
}

func (m *Infrastructure) EncodeConfig() {
	m.EncodedConfig = base64.StdEncoding.EncodeToString(utils.AesEny([]byte(*m.Config)))
}

type InfraTypes int8

const (
	InfraTypeAll      InfraTypes = 0
	InfraTypeK8s      InfraTypes = 1 // K8s配置
	InfraTypeRegistry InfraTypes = 2 // 镜像Registry
	InfraTypeAgent    InfraTypes = 3 // go-to-cloud代理
)

func GetInfra(orgID uint) ([]Infrastructure, error) {
	return getInfrastructures(orgID, InfraTypeAll)
}

func GetRegistries(orgID uint) ([]Infrastructure, error) {
	return getInfrastructures(orgID, InfraTypeRegistry)
}

func GetK8s(orgID uint) ([]Infrastructure, error) {
	return getInfrastructures(orgID, InfraTypeK8s)
}

func GetAgents(orgID uint) ([]Infrastructure, error) {
	return getInfrastructures(orgID, InfraTypeAgent)
}

// getInfrastructures 获取指定类型的基础设施
// orgId：所属组织
// infraType：基础设施类型；1：k8s；2：registry；3: 代理 0：所有
func getInfrastructures(orgId uint, infraType InfraTypes) ([]Infrastructure, error) {
	db := conf.GetDbClient()
	var org Org
	var err error

	if infraType > 0 {
		err = db.Preload("Infrastructures", "Type = ?", infraType).First(&org, orgId).Error
	} else {
		err = db.Preload("Infrastructures").First(&org, orgId).Error
	}
	return org.Infrastructures, err
}
