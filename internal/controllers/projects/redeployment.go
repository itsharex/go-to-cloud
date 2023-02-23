package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
)

// Redeployment
// @Tags Projects
// @Description 重新部署
// @Summary 重新部署
// @Param   ContentBody     body     deploy.Redeployment     true  "Request"     example(deploy.Redeployment)
// @Router /api/projects/{projectId}/deploy/redeploy [put]
// @Security JWT
func Redeployment(ctx *gin.Context) {
	exists, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var req deploy.Redeployment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if err := project.Redeployment(&req); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx)
	}
}
