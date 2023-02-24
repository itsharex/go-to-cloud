package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// QueryNamespaces 获取部署环境的可用名字空间
// @Tags Projects
// @Description 获取部署环境的可用名字空间
// @Summary 获取部署环境的可用名字空间
// @Success 200 {array} deploy.Deployment
// @Router /api/projects/{projectId}/deploy/apps [get]
// @Security JWT
func QueryNamespaces(ctx *gin.Context) {
	exists, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	projectIdStr := ctx.Param("projectId")
	if projectId, err := strconv.ParseUint(projectIdStr, 10, 64); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	} else {
		ns, err := project.ListNamespacesByProject(uint(projectId))
		if err == nil {
			response.Success(ctx, ns)
		} else {
			msg := err.Error()
			response.Fail(ctx, http.StatusInternalServerError, &msg)
		}
		return
	}

}
