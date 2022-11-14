package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const AgentTicket = "TODO_CHANGE_THIS_TICKET"

func AgentAuthHandler(c *gin.Context) {
	ticket := c.GetHeader("ticket")

	if nil == bcrypt.CompareHashAndPassword([]byte(ticket), []byte(AgentTicket)) {
		c.Next()
	}
}
