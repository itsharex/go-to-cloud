package configure

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/pkg/response"
)

// GetCodeRepos
// @BasePath /api
// @Tags Configure
// @Description 配置
// @Success 200
// @Router /api/configure/coderepo [get]
// @Security JWT
func GetCodeRepos(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"name":   "Hello",
		"avatar": "https://i.jd.com/defaultImgs/9.jpg",
	})
}
