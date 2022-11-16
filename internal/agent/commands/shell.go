package commands

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/agent/models"
	"go-to-cloud/internal/pkg/response"
)

// Shell
// @Tags Agent
// @Description 执行Shell命令
// @Summary 执行Shell命令
// @Param   ContentBody     body     models.ShellModel     true  "Request"     example(models.ShellModel)
// @Success 200
// @Router /commands/shell [POST]
// @Security JWT
func Shell(ctx *gin.Context) {
	var m models.ShellModel
	var err error
	if err = ctx.ShouldBind(&m); err != nil {
		response.BadRequest(ctx)
		return
	}

	// TODO: 执行shell命令
}
