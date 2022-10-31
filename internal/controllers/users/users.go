package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/util"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// Info
// @Tags User
// @Description 查看用户信息
// @Success 200
// @Router /api/user/info [get]
// @Security JWT
func Info(ctx *gin.Context) {
	exists, userId, userName, _, orgs := util.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	response.Success(ctx, gin.H{
		"userId":   userId,
		"userName": userName,
		"orgs":     orgs,
	})
}

// Logout
// @Tags User
// @Description 注销登录
// @Success 200
// @Router /api/user/logout [post]
// @Security JWT
func Logout(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"code": 20000,
		"data": gin.H{
			"name":   "Hello",
			"avatar": "https://i.jd.com/defaultImgs/9.jpg",
		},
	})
}
