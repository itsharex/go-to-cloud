package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/pkg/response"
)

// Info
// @Tags User
// @Description 查看用户信息
// @Success 200
// @Router /api/user/info [get]
// @Security JWT
func Info(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"name":   "Hello",
		"avatar": "https://i.jd.com/defaultImgs/9.jpg",
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
