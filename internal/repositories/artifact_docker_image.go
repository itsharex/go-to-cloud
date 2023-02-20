package repositories

import (
	"fmt"
	"go-to-cloud/conf"
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
	err := tx.Find(&images).Error

	return images, err
}

func CreateArtifact(image *ArtifactDockerImages) {
	db := conf.GetDbClient()

	_ = db.Model(&ArtifactDockerImages{}).Create(image)
}
