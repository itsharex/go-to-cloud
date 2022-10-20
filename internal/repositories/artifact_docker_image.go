package repositories

import (
	"go-to-cloud/conf"
)

type ArtifactDockerImages struct {
	Model
	Name           string `json:"name" gorm:"column:name"`
	ArtifactRepoID uint   `json:"artifactRepoId" gorm:"column:artifact_repo_id;index:artifact_docker_images_artifact_repo_id_index"`
	Tag            string `json:"tag" gorm:"column:tag;uniqueIndex:artifact_docker_images_hash_tag_uindex"`
	Hash           string `json:"hash" gorm:"column:hash;uniqueIndex:artifact_docker_images_hash_tag_uindex"`
	FullAddress    string `json:"fullAddress" gorm:"column:full_address"`
}

func (m *ArtifactDockerImages) TableName() string {
	return "artifact_docker_images"
}

// QueryImages 获取镜像仓库里的镜像列表
func QueryImages(artifactId uint) ([]ArtifactDockerImages, error) {
	db := conf.GetDbClient()

	var images []ArtifactDockerImages

	// TODO: 环境变量
	tx := db.Debug()

	tx = tx.Where(ArtifactDockerImages{ArtifactRepoID: artifactId})
	err := tx.Find(&images).Error

	return images, err
}
