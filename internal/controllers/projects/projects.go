package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/util"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
)

// List
// @Tags Projects
// @Description 查看项目信息
// @Success 200 {array} project.DataModel
// @Router /api/projects/list [get]
// @Security JWT
func List(ctx *gin.Context) {
	exists, _, _, orgs, _ := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	m, err := project.List(orgs)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, m)
	}
}

// CodeRepo
// @Tags Projects
// @Description 列出当前账户已绑定的SCM平台及可见的代码仓库
// @Success 200 {array} project.CodeRepoGroup
// @Router /api/projects/coderepo [get]
// @Security JWT
func CodeRepo(ctx *gin.Context) {
	exists, _, _, orgId, _ := util.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
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
