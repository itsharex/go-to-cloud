package project

import (
	project2 "go-to-cloud/internal/models/project"
	"go-to-cloud/internal/repositories"
)

// CreateNewProject 创建新项目
func CreateNewProject(userId uint, orgs []uint, model project2.DataModel) (uint, error) {
	return repositories.CreateProject(userId, orgs, model)
}
