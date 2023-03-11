package repositories

import (
	"fmt"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/artifact"
)

type ArtifactDockerImages struct {
	Model
	PipelineId     uint   `json:"pipelineId" gorm:"column:pipeline_id;type:bigint unsigned"`
	BuildId        uint   `json:"buildId" gorm:"column:build_id;type:bigint unsigned"`
	Name           string `json:"name" gorm:"column:name;type:varchar(200)"`
	ArtifactRepoID uint   `json:"artifactRepoId" gorm:"column:artifact_repo_id;index:artifact_docker_images_artifact_repo_id_index"`
	Tag            string `json:"tag" gorm:"column:tag;type:varchar(100)"`
	FullAddress    string `json:"fullAddress" gorm:"column:full_address;type:varchar(200)"`
}

// GetHashedCode 获取镜像唯一名称
func (m *ArtifactDockerImages) GetHashedCode() string {
	return fmt.Sprintf("%d,%d,%s", m.ArtifactRepoID, m.PipelineId, m.Name)
}

func (m *ArtifactDockerImages) TableName() string {
	return "artifact_docker_images"
}

// QueryImages 获取镜像仓库里的镜像列表
func QueryImages(artifactId uint) ([]ArtifactDockerImages, error) {
	db := conf.GetDbClient()

	var images []ArtifactDockerImages

	tx := db.Where(ArtifactDockerImages{ArtifactRepoID: artifactId})
	err := tx.Order("created_at DESC").Find(&images).Error

	return images, err
}

func CreateArtifact(image *ArtifactDockerImages) {
	db := conf.GetDbClient()

	_ = db.Model(&ArtifactDockerImages{}).Create(image)
}

func DeleteImages(userId, pipelineId, artifactRepoId uint) error {
	// TODO: 校验当前userId是否拥有数据删除权限

	tx := conf.GetDbClient()

	return tx.Where("pipeline_id = ? AND artifact_repo_id = ?", pipelineId, artifactRepoId).Delete(&ArtifactDockerImages{}).Error
}

func DeleteImage(userId, imageId uint) error {

	tx := conf.GetDbClient()

	// TODO: 校验当前userId是否拥有数据删除权限

	err := tx.Delete(&ArtifactDockerImages{
		Model: Model{
			ID: imageId,
		},
	}).Error

	return err
}

func QueryLatestImagesByProjectId(projectId uint) ([]artifact.FullName, error) {
	db := conf.GetDbClient()

	var images []artifact.FullName

	tx := db.Raw(`
select a.id, a.name, a.tag, a.full_address address
from artifact_docker_images a
         inner join pipeline p on p.id = a.pipeline_id
         inner join (select max(d.updated_at) upd, d.pipeline_id pipelineId
                     from artifact_docker_images d
                     group by d.pipeline_id) x
                    on x.upd = a.updated_at AND x.pipelineId = a.pipeline_id
where p.project_id = ?`, projectId).Find(&images)

	return returnWithError(images, tx.Error)
}

func QueryImageTagsById(id uint) ([]string, error) {
	db := conf.GetDbClient()

	var tags []string

	tx := db.Raw(`
	select tag
		from artifact_docker_images b
		inner join (select pipeline_id
			from artifact_docker_images a
			where a.id = ?) x on x.pipeline_id = b.pipeline_id
			order by b.id desc`, id).Find(&tags)

	return returnWithError(tags, tx.Error)
}
