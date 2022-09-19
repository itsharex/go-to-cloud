package auth

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/middlewares"
)

// Login
// @BasePath /
// @Tags User
// @Description 查看用户信息
// @Accept json
// @Product json
// @Param login body models.LoginModel true "Login Model"
// @Router /login [post]
// @Success 200
func Login(ctx *gin.Context) {
	middlewares.GinJwtMiddleware().LoginHandler(ctx)
}
