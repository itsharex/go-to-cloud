package commands

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/agent/models"
	"go-to-cloud/internal/pkg/pipeline/stages"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// Clone
// @Tags Agent
// @Description 克隆代码
// @Summary 克隆代码
// @Param   ContentBody     body     models.GitModel     true  "Request"     example(models.GitModel)
// @Success 200
// @Router /commands/clone [POST]
// @Security JWT
func Clone(c *gin.Context) {
	var m models.GitModel
	if err := c.ShouldBind(&m); err != nil {
		response.BadRequest(c)
		return
	}

	// TODO: 默认工作目录
	workdir := ""
	if err := stages.GitClone(&m.Address, &m.Branch, &workdir, m.DecodeToken()); err != nil {
		response.Fail(c, http.StatusInternalServerError, nil, err)
	} else {
		response.Success(c)
	}
}
