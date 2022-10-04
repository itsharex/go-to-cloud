package scm

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/util"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/pkg/scm"
	"net/http"
)

// GetCodeRepos
// @Tags Configure
// @Description 代码仓库配置
// @Success 200
// @Router /api/configure/coderepo [get]
// @Security JWT
func GetCodeRepos(ctx *gin.Context) {
	exists, userId, userName, orgs := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, "unauthorized")
		return
	}

	response.Success(ctx, gin.H{
		"userId":   userId,
		"userName": userName,
		"orgs":     orgs,
	})
}

// BindCodeRepo 绑定代码仓库
// @Tags Configure
// @Description 代码仓库配置
// @Success 200
// @Router /api/configure/coderepo/bind [post]
// @Security JWT
func BindCodeRepo(ctx *gin.Context) {
	var req models.Scm
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	success, err := scm.Ping(&req.ScmTesting)
	if err != nil {
		response.Fail(ctx, http.StatusForbidden, err.Error())
		return
	}
	if !success {
		response.Fail(ctx, http.StatusForbidden, "scm connection failed")
		return
	}

	exists, userId, _, orgs := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, "unauthorized")
		return
	}

	err = scm.Bind(&req, userId, orgs)

	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
