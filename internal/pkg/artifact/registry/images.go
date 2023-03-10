package registry

import (
	"go-to-cloud/internal/models/artifact/docker_image"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// extractLatestVer 计算最新版本，将src中的发布时间与dst中的发布时间对比，如果更新，则替换dst中的latestVer
func extractLatestVer(dst *docker_image.Image, src *repositories.ArtifactDockerImages) {
	if dst.LatestPublishAt.Before(src.CreatedAt) {
		dst.LatestVer = src.Tag
		dst.LatestPublishAt = src.CreatedAt
	}
}

func QueryImages(artifactID uint) ([]*docker_image.Image, error) {
	images, err := repositories.QueryImages(artifactID)

	if err != nil {
		return nil, err
	}

	var rlt []*docker_image.Image
	if len(images) > 0 {
		aggImage := make(map[string]*docker_image.Image)
		for _, image := range images {
			hashedCode := image.GetHashedCode()
			if aggImage[hashedCode] == nil {
				aggImage[hashedCode] = &docker_image.Image{
					Hash:            image.GetHashedCode(),
					Name:            image.Name,
					FullName:        image.FullAddress,
					Tags:            make([]docker_image.Tag, 0),
					PublishedAt:     utils.JsonTime(image.CreatedAt),
					LatestVer:       image.Tag,
					LatestPublishAt: image.CreatedAt,
				}
			}
			extractLatestVer(aggImage[hashedCode], &image)
			aggImage[hashedCode].Tags = append(aggImage[hashedCode].Tags, docker_image.Tag{
				Tag:         image.Tag,
				PublishedAt: utils.JsonTime(image.CreatedAt),
			})
		}
		rlt = make([]*docker_image.Image, len(aggImage))
		i := 0
		for _, image := range aggImage {
			rlt[i] = image
			i++
		}
	}

	return rlt, nil
}
