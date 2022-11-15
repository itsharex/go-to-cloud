package commands

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/middlewares"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// GenToken
// @Tags Agent
// @Description 获取Token
// @Summary 获取Token
// @Router /commands/dev/token [get]
// @Success 200
func GenToken(c *gin.Context) {
	hashBytes, _ := bcrypt.GenerateFromPassword([]byte(middlewares.AgentTicket), bcrypt.DefaultCost)
	c.JSON(http.StatusOK, string(hashBytes))
}
