package commands

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/agent/models"
	"go-to-cloud/internal/pkg/pipeline/stages"
	"go-to-cloud/internal/pkg/response"
	"net/http"
	"os"
)

// Clone
// @Tags Agent
// @Description 克隆代码
// @Summary 克隆代码
// @Param   ContentBody     body     models.GitModel     true  "Request"     example(models.GitModel)
// @Success 200 {string} workdir
// @Router /commands/clone [POST]
// @Security JWT
func Clone(c *gin.Context) {
	var m models.GitModel
	var err error
	if err = c.ShouldBind(&m); err != nil {
		response.BadRequest(c)
		return
	}

	var workdir string
	if workdir, err = os.MkdirTemp("", "gtc"); err != nil {
		response.Fail(c, http.StatusInternalServerError, nil, err)
		return
	}

	gitCloneStage := stages.GitCloneStage{
		Token:   *m.DecodeToken(),
		GitUrl:  m.Address,
		Branch:  m.Branch,
		WorkDir: workdir,
	}
	if err = gitCloneStage.Run(); err != nil {
		response.Fail(c, http.StatusInternalServerError, nil, err)
	} else {
		response.Success(c, workdir)
	}
}
