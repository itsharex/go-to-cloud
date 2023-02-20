package repositories

import (
	"go-to-cloud/conf"
)

type ArtifactDockerImages struct {
	Model
	PipelineId     uint   `json:"pipelineId" gorm:"column:pipeline_id;type:bigint"`
	Name           string `json:"name" gorm:"column:name"`
	ArtifactRepoID uint   `json:"artifactRepoId" gorm:"column:artifact_repo_id;index:artifact_docker_images_artifact_repo_id_index"`
	Tag            string `json:"tag" gorm:"column:tag;type:text;"`
	Hash           string `json:"hash" gorm:"column:hash;"`
	FullAddress    string `json:"fullAddress" gorm:"column:full_address"`
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
