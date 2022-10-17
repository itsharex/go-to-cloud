package k8s

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/util"
	k8smodel "go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/pkg/deploy/k8s"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// UpdateK8sRepo 更新K8s仓库
// @Tags Configure
// @Description k8s仓库配置
// @Success 200
// @Param   ContentBody     body     k8s.K8s     true  "Request"     example(k8s.K8s)
// @Router /api/configure/deploy/k8s [put]
// @Security JWT
func UpdateK8sRepo(ctx *gin.Context) {
	var req k8smodel.K8s
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	success, err := k8s.Ping(&req.Testing)
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

	err = k8s.Update(&req, userId, orgs)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
