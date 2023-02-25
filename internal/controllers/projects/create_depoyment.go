package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// CreateDeployment 创建部署应用
// @Tags Projects
// @Description 创建部署应用
// @Summary 创建部署应用
// @Param   ContentBody     body     deploy.Deployment     true  "Request"     example(deploy.Deployment)
// @Router /api/projects/{projectId}/deploy/app [post]
// @Security JWT
func CreateDeployment(ctx *gin.Context) {
	exists, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	projectIdStr := ctx.Param("projectId")
	projectId, _ := strconv.ParseUint(projectIdStr, 10, 64)

	var req deploy.Deployment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if err := project.CreateDeployments(uint(projectId), &req); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, gin.H{
			"success": true,
		})
	}
}
