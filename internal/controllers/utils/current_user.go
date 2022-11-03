package utils

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CurrentUser(ctx *gin.Context) (exists bool, userId uint, user *string, orgIds []uint, orgs map[uint]string) {

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
			orgs = make(map[uint]string, sz)
			orgIds = make([]uint, 0)
			for key, val := range maps {
				orgId, _ := strconv.ParseUint(key, 10, 64)
				orgs[uint(orgId)] = val.(string)
				orgIds = append(orgIds, uint(orgId))
			}
		}
	}
	userId = uint(jti)
	user = &sub

	exists = true

	return
}
