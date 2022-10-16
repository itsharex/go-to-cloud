package registry

import (
	"go-to-cloud/internal/models/artifact/docker_image"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func QueryImages(artifactID uint) ([]*docker_image.Image, error) {
	images, err := repositories.QueryImages(artifactID)

	if err != nil {
		return nil, err
	}

	var rlt []*docker_image.Image
	if len(images) > 0 {
		aggImage := make(map[string]*docker_image.Image)
		for _, image := range images {
			if aggImage[image.Hash] == nil {
				aggImage[image.Hash] = &docker_image.Image{
					Name:        image.Name,
					FullName:    image.FullAddress,
					Tags:        make([]string, 0),
					PublishedAt: utils.JsonTime(image.CreatedAt),
				}
			}
			aggImage[image.Hash].Tags = append(aggImage[image.Hash].Tags, image.Tag)
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
