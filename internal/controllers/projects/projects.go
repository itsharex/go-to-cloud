package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/util"
	projectModel "go-to-cloud/internal/models/project"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
)

// List
// @Tags Projects
// @Description 查看项目信息
// @Success 200
// @Router /api/projects/list [get]
// @Security JWT
func List(ctx *gin.Context) {
	m := make([]projectModel.DataModel, 0)
	m = append(m, projectModel.DataModel{
		Id:   0,
		Name: "aaa",
	})
	m = append(m, projectModel.DataModel{
		Id:   1,
		Name: "bbb",
	})
	response.Success(ctx, m)
}

// CodeRepo
// @Tags Projects
// @Description 列出当前账户已绑定的SCM平台及可见的代码仓库
// @Success 200 {array} project.CodeRepoGroup
// @Router /api/projects/coderepo [get]
// @Security JWT
func CodeRepo(ctx *gin.Context) {
	exists, _, _, orgs := util.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	i := 0
	orgId := make([]uint, len(orgs))
	for k, _ := range orgs {
		orgId[i] = k
		i++
	}

	m, err := project.GetCodeRepoGroupsByOrg(orgId)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, m)
		return
	}
}
