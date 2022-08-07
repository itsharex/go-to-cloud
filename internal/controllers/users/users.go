package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/pkg/response"
)

// Login 一个跑通全流程的示例，业务代码未补充完整
// @BasePath /api
// Login 登录 godoc
// @Summary 用户登录
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Login
// @Router /api/user/Login [get]
func Login(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"username":    "admin",
		"password":    "admin",
		"role":        "admin",
		"roleId":      "1",
		"permissions": []string{"*.*.*"},
	})
}

func Info(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"name":   "Hello",
		"avatar": "https://i.jd.com/defaultImgs/9.jpg",
	})
}

func Logout(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"code": 20000,
		"data": gin.H{
			"name":   "Hello",
			"avatar": "https://i.jd.com/defaultImgs/9.jpg",
		},
	})
}
