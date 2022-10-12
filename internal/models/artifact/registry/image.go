package registry

type Image struct {
	Name     string      // 镜像名
	FullName string      // 完整路径，包含项目名称，e.g. library/mysql
	Tags     []TagDetail // Tag列表，按从旧到新的顺序
}

type TagDetail struct {
	Ver string // tag ver
}
