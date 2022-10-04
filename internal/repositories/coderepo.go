package repositories

import (
	"go-to-cloud/conf"
	"go-to-cloud/internal/models"
	"gorm.io/gorm"
	"time"
)

type Coderepo struct {
	gorm.Model
	ScmOrigin   int    `json:"scm_origin" gorm:"column:scm_origin"`     // 代码仓库来源；gitlab(0);github(1);gitee(2);gitea(3)
	IsPublic    int8   `json:"is_public" gorm:"column:is_public"`       // 是否公开仓库
	AccessToken string `json:"access_token" gorm:"column:access_token"` // 访问令牌 PAT
	Url         string `json:"url" gorm:"column:url"`                   // SCM平台地址（非项目仓库地址）
	CreatedBy   int64  `json:"created_by" gorm:"column:created_by"`     // 	仓库创建人
	OwnedOrgBy  int64  `json:"owned_org_by" gorm:"column:owned_org_by"` // SCM所属组织
	Orgs        *Org   `gorm:"foreignKey:OwnedOrgBy"`
	Remark      string `json:"remark" gorm:"column:remark"'`
}

func (m *Coderepo) TableName() string {
	return "coderepo"
}

// BindCodeRepo 绑定代码仓库
func BindCodeRepo(model *models.Scm, userId int64, orgs []int64) error {
	repos := make([]Coderepo, len(orgs))
	for i, orgId := range orgs {
		isPublic := int8(0)
		if model.IsPublic {
			isPublic = 1
		} else {
			isPublic = 0
		}
		repos[i] = Coderepo{
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			ScmOrigin:   int(model.Origin),
			IsPublic:    isPublic,
			AccessToken: *model.Token,
			Url:         model.Url,
			CreatedBy:   userId,
			OwnedOrgBy:  orgId,
			Remark:      model.Remark,
		}
	}

	return conf.GetDbClient().CreateInBatches(repos, len(repos)).Error
}
