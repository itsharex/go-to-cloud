package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/pkg/response"
)

// List
// @BasePath /api
// @Tags Projects
// @Description 查看项目信息
// @Success 200
// @Router /api/projects/list [get]
// @Security JWT
func List(ctx *gin.Context) {
	m := make([]models.ProjectDataModel, 0)
	m = append(m, models.ProjectDataModel{
		Id:   0,
		Name: "aaa",
	})
	m = append(m, models.ProjectDataModel{
		Id:   1,
		Name: "bbb",
	})
	response.Success(ctx, gin.H{
		"data": m,
	})
}
