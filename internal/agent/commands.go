package agent

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Healthz 健康检查
// @Tags Agent
// @Description 健康检查
// @Router /commands/healthz [head]
// @Success 204
func Healthz(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})
}
