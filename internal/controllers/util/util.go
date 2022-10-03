package util

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func CurrentUser(ctx *gin.Context) (userId int64, user *string) {
	mapping := ctx.MustGet("JWT_PAYLOAD").(jwt.MapClaims)

	jti := mapping["jti"].(float64)
	sub := mapping["sub"].(string)
	orgs := mapping["orgs"]

	_ = orgs
	userId = int64(jti)
	user = &sub
	return
}
