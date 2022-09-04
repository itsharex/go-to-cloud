package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Index 首页
// @BasePath /index
// @Schemes
// @Description
// @Tags Index
// @Accept json
// @Success 302 {string} Index
// @Router /index [get]
func Index(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/user/info")
	ctx.AbortWithStatus(http.StatusFound)
}
