package index

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/pkg/response"
)

// Jump 跳转引导
// @BasePath /index
// Jump 跳转引导
// @Summary 跳转引导
// @Schemes
// @Description 跳转引导，如果首次登录，则跳转至init，如果未登录，则跳转至login，否则进入home
// @Tags Index
// @Accept json
// @Success 302 {string} Jump
// @Router /index/jump [get]
func Jump(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"username":    "admin",
		"password":    "admin",
		"role":        "admin",
		"roleId":      "1",
		"permissions": []string{"*.*.*"},
	})
}
