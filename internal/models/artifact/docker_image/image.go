package docker_image

import "go-to-cloud/internal/utils"

type Image struct {
	Name        string   // 镜像名
	FullName    string   // 完整路径，包含项目名称，e.g. library/mysql
	Tags        []string // Tag
	PublishedAt utils.JsonTime
}
