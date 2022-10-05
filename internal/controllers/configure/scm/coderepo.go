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
	exists, userId, _, orgs := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var query models.ScmQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	orgsId := make([]int64, len(orgs))
	idx := int64(0)
	for key, _ := range orgs {
		orgsId[idx] = key
		idx++
	}
	result, err := scm.List(userId, orgsId, &query)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, result)
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
		msg := err.Error()
		response.Fail(ctx, http.StatusForbidden, &msg)
		return
	}
	if !success {
		response.Fail(ctx, http.StatusForbidden, nil)
		return
	}

	exists, userId, _, orgs := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = scm.Bind(&req, userId, orgs)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
