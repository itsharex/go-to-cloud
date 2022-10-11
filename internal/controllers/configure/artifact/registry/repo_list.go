package registry

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/pkg/artifact/registry"
	"go-to-cloud/internal/pkg/response"
	"strconv"
)

// ListRepositories 从Registry获取镜像仓库列表
// @Tags Configure
// @Description 制品仓库配置
// @Success 200
// @Router /api/configure/artifact/registry/:id [get]
// @Security JWT
func ListRepositories(ctx *gin.Context) {
	val := ctx.Param("id")

	repoId, err := strconv.ParseUint(val, 10, 64)

	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	registry.ListRepositories(uint(repoId))
}
