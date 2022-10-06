package scm

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/util"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/pkg/scm"
	"net/http"
	"strconv"
)

// QueryCodeRepos
// @Tags Configure
// @Description 代码仓库配置
// @Success 200
// @Router /api/configure/coderepo [get]
// @Security JWT
func QueryCodeRepos(ctx *gin.Context) {
	exists, _, _, orgs := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var query models.ScmQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	orgsId := make([]uint, len(orgs))
	idx := uint(0)
	for key := range orgs {
		orgsId[idx] = key
		idx++
	}
	result, err := scm.List(orgsId, &query)

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

// UpdateCodeRepo 更新代码仓库
// @Tags Configure
// @Description 代码仓库配置
// @Success 200
// @Router /api/configure/coderepo/bind [put]
// @Security JWT
func UpdateCodeRepo(ctx *gin.Context) {
	var req models.Scm
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	success, err := scm.Ping(&req.ScmTesting)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusOK, &msg)
		return
	}
	if !success {
		response.Fail(ctx, http.StatusOK, nil)
		return
	}

	exists, userId, _, orgs := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = scm.Update(&req, userId, orgs)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}

// RemoveCodeRepo 移除代码仓库
// @Tags Configure
// @Description 代码仓库配置
// @Success 200
// @Router /api/configure/coderepo/:id [delete]
// @Security JWT
func RemoveCodeRepo(ctx *gin.Context) {
	val := ctx.Param("id")

	repoId, err := strconv.ParseUint(val, 10, 64)

	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	exists, userId, _, _ := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = scm.RemoveRepo(userId, uint(repoId))

	var message string
	if err != nil {
		message = err.Error()
	} else {
		message = ""
	}
	response.Success(ctx, gin.H{
		"success": err == nil,
		"message": message,
	})
}
