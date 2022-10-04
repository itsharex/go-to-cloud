package scm

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/pkg/scm"
	"net/http"
)

// Testing
// @Tags Configure
// @Description 代码仓库配置
// @Produce json
// @Accept json
// @Param   ContentBody     body     models.ScmTesting     true  "Request"     example(models.ScmTesting)
// @Security JWT
// @Success 200
// @Router /api/configure/coderepo/testing [post]
func Testing(ctx *gin.Context) {
	var req models.ScmTesting
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if success, err := scm.Ping(&req); err != nil {
		response.Fail(ctx, http.StatusForbidden, err.Error())
		return
	} else {
		response.Success(ctx, gin.H{
			"success": success,
		})
	}
}
