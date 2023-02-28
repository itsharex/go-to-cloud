package monitor

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/monitor"
	"net/http"
)

// List
// @Tags Monitor
// @Description 列出安装的应用
// @Success 200 {array} deploy.DeploymentDescription
// @Router /api/monitor/{k8s}/list/ [get]
// @Param        projectId    query     string  false  "project id"
// @Param        deploymentId    query     string  false  "deployment id"
// @Param        k8s    path     string  true  "k8s repo id"
// @Security JWT
func List(ctx *gin.Context) {
	exists, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	k8sRepoId, err := getUIntParamFromQueryOrPath("k8s", ctx, false)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}
	projectId, err := getUIntParamFromQueryOrPath("projectId", ctx, true)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}
	deploymentId, err := getUIntParamFromQueryOrPath("deploymentId", ctx, true)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	m, err := monitor.QueryApps(projectId, k8sRepoId, deploymentId)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, m)
	}
}
