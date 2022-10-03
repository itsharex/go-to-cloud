package scm

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/util"
	"go-to-cloud/internal/pkg/response"
	"net/http"
	"strconv"
)

// GetCodeRepos
// @Tags Configure
// @Description 代码仓库配置
// @Success 200
// @Router /api/configure/coderepo [get]
// @Security JWT
func GetCodeRepos(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"name":   "Hello",
		"avatar": "https://i.jd.com/defaultImgs/9.jpg",
	})
}

// CreateCodeRepo
// @Tags Configure
// @Description 代码仓库配置
// @Success 200
// @Router /api/configure/coderepo [post]
// @Security JWT
func CreateCodeRepo(ctx *gin.Context) {

	userId, user := util.CurrentUser(ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"jti": strconv.Itoa(int(userId)),
		"sub": *user,
	})
}
