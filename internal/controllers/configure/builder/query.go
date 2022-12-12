package builder

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/builder"
	builder2 "go-to-cloud/internal/pkg/builder"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// QueryNodesOnK8s
// @Tags Configure
// @Description 节点管理
// @Success 200 {array} builder.NodesOnK8s
// @Router /api/configure/builder/nodes/k8s [get]
// @Security JWT
func QueryNodesOnK8s(ctx *gin.Context) {
	exists, _, _, orgsId, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var query builder.Query
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	result, err := builder2.ListNodesOnK8s(orgsId, &query)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, result)
}
