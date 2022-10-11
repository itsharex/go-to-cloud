package artifact

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/util"
	artifactModels "go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/pkg/artifact"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// QueryArtifactRepos
// @Tags Configure
// @Description 制品仓库配置
// @Success 200 {object} artifact.Artifact
// @Router /api/configure/artifact [get]
// @Security JWT
func QueryArtifactRepos(ctx *gin.Context) {
	exists, _, _, orgs := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var query artifactModels.Query
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	orgsId := make([]uint, len(orgs))
	idx := uint(0)
	for key := range orgs {
		orgsId[idx] = key
		idx++
	}
	result, err := artifact.List(orgsId, &query)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, result)
}
