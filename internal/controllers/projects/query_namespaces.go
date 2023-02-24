package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
)

// QueryNamespaces 获取部署环境的可用名字空间
// @Tags Projects
// @Description 根据当前用户所属组织获取部署环境的可用名字空间
// @Summary 获取部署环境的可用名字空间
// @Success 200 {array} deploy.Deployment
// @Router /api/projects/{projectId}/deploy/namespaces [get]
// @Security JWT
func QueryNamespaces(ctx *gin.Context) {
	exists, _, _, orgs, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	ns, err := project.ListNamespacesByOrg(orgs)
	if err == nil {
		response.Success(ctx, ns)
	} else {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	}

}
