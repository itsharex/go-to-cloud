package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/models/project"
	"go-to-cloud/internal/pkg/response"
)

// List
// @Tags Projects
// @Description 查看项目信息
// @Success 200
// @Router /api/projects/list [get]
// @Security JWT
func List(ctx *gin.Context) {
	m := make([]project.DataModel, 0)
	m = append(m, project.DataModel{
		Id:   0,
		Name: "aaa",
	})
	m = append(m, project.DataModel{
		Id:   1,
		Name: "bbb",
	})
	response.Success(ctx, gin.H{
		"data": m,
	})
}

// CodeRepo
// @Tags Projects
// @Description 列出当前账户已绑定的SCM平台及可见的代码仓库
// @Success 200 {array} project.CodeRepoGroup
// @Router /api/projects/coderepo [get]
// @Security JWT
func CodeRepo(ctx *gin.Context) {
	m := make([]project.CodeRepoGroup, 0)
	m = append(m, project.CodeRepoGroup{
		Id:   0,
		Name: "aaa",
		Git: []project.GitSources{
			{
				Name: "Te1",
				Url:  "http://dsfsaf.git",
			},
		},
	})
	response.Success(ctx, gin.H{
		"data": m,
	})
}
