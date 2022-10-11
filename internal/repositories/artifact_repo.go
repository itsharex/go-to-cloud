package repositories

import (
	"encoding/json"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/artifact"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type ArtifactRepo struct {
	gorm.Model
	Name           string         `json:"name" gorm:"column:name"`
	ArtifactOrigin int            `json:"artifactOrigin" gorm:"column:artifact_origin"` // 制品仓库来源；Docker(1);
	IsSecurity     int8           `json:"isSecurity" gorm:"column:is_security"`         // 是否使用https
	Account        string         `json:"account" gorm:"column:account"`                // 用户名
	Password       string         `json:"password" gorm:"column:password"`              // 密码
	Url            string         `json:"url" gorm:"column:url"`                        // 制品仓库平台地址
	CreatedBy      uint           `json:"createdBy" gorm:"column:created_by"`           // 仓库创建人
	BelongsTo      datatypes.JSON `json:"belongsTo" gorm:"column:belongs_to;"`          // SCM所属组织
	Remark         string         `json:"remark" gorm:"column:remark"`
}

func (m *ArtifactRepo) TableName() string {
	return "artifact_repo"
}

type MergedArtifactRepoWithOrg struct {
	ArtifactRepo
	Org []OrgLite
}

type ArtifactRepoWithOrg struct {
	ArtifactRepo
	OrgLite
}

func mergeArtifactRepoOrg(repos []ArtifactRepoWithOrg) ([]MergedArtifactRepoWithOrg, error) {
	r := make(map[uint][]OrgLite)
	for _, repo := range repos {
		x := r[repo.ID]
		if x == nil {
			r[repo.ID] = make([]OrgLite, 0)
		}
		r[repo.ID] = append(r[repo.ID], OrgLite{
			OrgId:   repo.OrgId,
			OrgName: repo.OrgName,
		})
	}

	merged := make(map[uint]*MergedArtifactRepoWithOrg)
	for _, repo := range repos {
		if merged[repo.ID] == nil {
			merged[repo.ID] = &MergedArtifactRepoWithOrg{
				ArtifactRepo: ArtifactRepo{
					Model: gorm.Model{
						ID:        repo.ID,
						CreatedAt: repo.CreatedAt,
						UpdatedAt: repo.UpdatedAt,
						DeletedAt: repo.DeletedAt,
					},
					Name:           repo.Name,
					ArtifactOrigin: repo.ArtifactOrigin,
					IsSecurity:     repo.IsSecurity,
					Account:        repo.Account,
					Password:       repo.Password,
					Url:            repo.Url,
					CreatedBy:      repo.CreatedBy,
					BelongsTo:      datatypes.JSON{},
					Remark:         repo.Remark,
				},
				Org: r[repo.ID],
			}
		}
	}
	rlt := make([]MergedArtifactRepoWithOrg, len(merged))
	counter := 0
	for _, m := range merged {
		rlt[counter] = *m
		counter++
	}
	return rlt, nil
}

func buildArtifactRepo(model *artifact.Artifact, userId uint, orgs []uint, gormModel *gorm.Model) (*ArtifactRepo, error) {
	isSecurity := int8(0)
	if model.IsSecurity {
		isSecurity = 1
	} else {
		isSecurity = 0
	}
	belongs, err := json.Marshal(orgs)
	if err != nil {
		return nil, err
	}
	repo := ArtifactRepo{
		Model:          *gormModel,
		ArtifactOrigin: int(model.Type),
		Name:           model.Name,
		IsSecurity:     isSecurity,
		Account:        model.User,
		Password:       model.Password,
		Url:            model.Url,
		CreatedBy:      userId,
		BelongsTo:      datatypes.JSON(belongs),
		Remark:         model.Remark,
	}

	return &repo, nil
}

// BindArtifactRepo 绑定制品仓库
func BindArtifactRepo(model *artifact.Artifact, userId uint, orgs []uint) error {
	g := &gorm.Model{
		CreatedAt: time.Now(),
	}
	repo, err := buildArtifactRepo(model, userId, orgs, g)
	if err != nil {
		return err
	}

	tx := conf.GetDbClient()

	// TODO: os.Env
	tx = tx.Debug()

	err = tx.Omit("updated_at").Create(&repo).Error
	return err
}

// QueryArtifactRepo 查询制品仓库
func QueryArtifactRepo(orgs []uint, repoNamePattern string) ([]MergedArtifactRepoWithOrg, error) {
	var repo []ArtifactRepoWithOrg

	tx := conf.GetDbClient().Model(&ArtifactRepo{})

	tx = tx.Select("artifact_repo.*, org.Id AS orgId, org.Name AS orgName")
	tx = tx.Joins("INNER JOIN org ON JSON_CONTAINS(artifact_repo.belongs_to, cast(org.id as JSON), '$')")
	tx = tx.Where("org.ID IN ? AND org.deleted_at IS NULL", orgs)

	if len(repoNamePattern) > 0 {
		tx = tx.Where("artifact_repo.name like ?", repoNamePattern+"%")
	}

	tx = tx.Order("created_at desc")

	// TODO: os.Env
	tx = tx.Debug()

	err := tx.Scan(&repo).Error

	if err == nil {
		return mergeArtifactRepoOrg(repo)
	} else {
		return nil, err
	}
}
