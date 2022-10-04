package util

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CurrentUser(ctx *gin.Context) (exists bool, userId int64, user *string, orgs map[int64]string) {

	defer func() {
		if r := recover(); r != nil {
			exists = false
		}
	}()

	mapping := ctx.MustGet("JWT_PAYLOAD").(jwt.MapClaims)

	jti := mapping["jti"].(float64)
	sub := mapping["sub"].(string)
	orgsMaps := mapping["orgs"]

	if orgsMaps != nil {
		maps := orgsMaps.(map[string]interface{})
		if sz := len(maps); sz > 0 {
			orgs = make(map[int64]string, sz)
			for key, val := range maps {
				orgId, _ := strconv.ParseInt(key, 10, 64)
				orgs[orgId] = val.(string)
			}
		}
	}
	userId = int64(jti)
	user = &sub

	exists = true

	return
}
