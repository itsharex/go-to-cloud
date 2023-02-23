package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
)

// Scale
// @Tags Projects
// @Description 伸缩容器
// @Summary 伸缩容器
// @Param   ContentBody     body     deploy.ScalePods     true  "Request"     example(deploy.ScalePods)
// @Router /api/projects/{projectId}/deploy/scale [put]
// @Security JWT
func Scale(ctx *gin.Context) {
	exists, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var req deploy.ScalePods
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if err := project.ScaleDeployment(&req); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx)
	}
}
