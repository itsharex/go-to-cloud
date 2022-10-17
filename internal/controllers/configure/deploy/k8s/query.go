package k8s

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/util"
	scm2 "go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/pkg/scm"
	"net/http"
)

// QueryK8sRepos
// @Tags Configure
// @Description k8s环境配置
// @Success 200 {object} scm.Scm
// @Router /api/configure/deploy/k8s [get]
// @Security JWT
func QueryK8sRepos(ctx *gin.Context) {
	exists, _, _, orgs := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var query scm2.Query
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
