package repositories

import (
	"crypto/md5"
	"errors"
	"fmt"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models"
	"gorm.io/gorm"
	"io"
	"strconv"
	"strings"
	"time"
)

type CodeRepo struct {
	gorm.Model
	Name          string `json:"name" gorm:"column:name"`
	ScmOrigin     int    `json:"scm_origin" gorm:"column:scm_origin"`     // 代码仓库来源；gitlab(0);github(1);gitee(2);gitea(3)
	IsPublic      int8   `json:"is_public" gorm:"column:is_public"`       // 是否公开仓库
	AccessToken   string `json:"access_token" gorm:"column:access_token"` // 访问令牌 PAT
	Url           string `json:"url" gorm:"column:url"`                   // SCM平台地址（非项目仓库地址）
	CreatedBy     int64  `json:"created_by" gorm:"column:created_by"`     // 仓库创建人
	OwnedOrgBy    int64  `json:"owned_org_by" gorm:"column:owned_org_by"` // SCM所属组织
	Orgs          []Org  `json:"orgs" gorm:"-"`
	Remark        string `json:"remark" gorm:"column:remark"`
	FingerPrinter string `json:"fingerPrinter" gorm:"column:finger_printer"` // 仓库指纹
}

func (m *CodeRepo) TableName() string {
	return "coderepo"
}

type MultiInt64 []int64

func (MultiInt64) GormDataType() string {
	return "text"
}

func (s *MultiInt64) Scan(src interface{}) error {
	str, ok := src.([]uint8)
	if !ok {
		return errors.New("failed to scan MultiString field - source is not a []uint8")
	}
	t := strings.Split(string(str), ",")
	x := make([]int64, len(t))
	for i, s2 := range t {
		x[i], _ = strconv.ParseInt(s2, 10, 64)
	}
	*s = x
	return nil
}

type MergedCodeRepo struct {
	CodeRepo
	OrgsId MultiInt64 `gorm:"column:orgsId; type:text"`
}

func (m *CodeRepo) SetFinerPrinter(fingerPrinter *string) {
	if fingerPrinter != nil {
		m.FingerPrinter = *fingerPrinter
	}
}

// ExtractFingerPrinter 提取仓库指纹
func ExtractFingerPrinter(repo *CodeRepo) *string {
	raw := fmt.Sprintf("%d:%d:%d:%s:%s", repo.CreatedBy, repo.ScmOrigin, repo.IsPublic, repo.AccessToken, repo.Url)
	h := md5.New()
	io.WriteString(h, raw)
	hashed := fmt.Sprintf("%x", h.Sum(nil))
	return &hashed
}

func ReadCodeRepo(orgs []int64, repoNamePattern string) ([]MergedCodeRepo, error) {
	var repo []MergedCodeRepo

	tx := conf.GetDbClient().Where("owned_org_by IN ?", orgs)

	if len(repoNamePattern) > 0 {
		tx = tx.Where("name LIKE ?", repoNamePattern+"%")
	}
	tx = tx.Debug()

	if conf.GetDbClient().Dialector.Name() == "mysql" {
		tx = tx.Exec(*disableOnlyFullGroupBy())
	}

	tx = tx.Select("*, GROUP_CONCAT(owned_org_by) AS orgsId")

	err := tx.Group("finger_printer").Find(&repo).Error

	if err == nil {
		for i, codeRepo := range repo {
			//codeRepo.Orgs = make([]Org, 0)
			conf.GetDbClient().Debug().Find(&repo[i].Orgs, codeRepo.OrgsId)
		}
	}
	return repo, err
}

// BindCodeRepo 绑定代码仓库
func BindCodeRepo(model *models.Scm, userId int64, orgs []int64) error {
	repos := make([]CodeRepo, len(orgs))
	for i, orgId := range orgs {
		isPublic := int8(0)
		if model.IsPublic {
			isPublic = 1
		} else {
			isPublic = 0
		}
		repos[i] = CodeRepo{
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			ScmOrigin:   int(model.Origin),
			Name:        model.Name,
			IsPublic:    isPublic,
			AccessToken: *model.Token,
			Url:         model.Url,
			CreatedBy:   userId,
			OwnedOrgBy:  orgId,
			Remark:      model.Remark,
		}
		repos[i].SetFinerPrinter(ExtractFingerPrinter(&repos[i]))
	}

	return conf.GetDbClient().CreateInBatches(repos, len(repos)).Error
}
