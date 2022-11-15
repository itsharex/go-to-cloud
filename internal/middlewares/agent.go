package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const AgentTicket = "TODO_CHANGE_THIS_TICKET"

func AgentAuthHandler(c *gin.Context) {
	ticket := c.GetHeader("Authorization")

	if nil == bcrypt.CompareHashAndPassword([]byte(ticket), []byte(AgentTicket)) {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, "incorrect ticket")
	}
}
