package monitor

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getUIntParamFromQueryOrPath(paramKey string, ctx *gin.Context, allowNull bool) (uint, error) {

	keyStr := ctx.Param(paramKey)
	if allowNull {
		if len(keyStr) == 0 {
			return 0, nil
		} else {
			return 0, errors.New("argument '" + paramKey + "' empty or null")
		}
	}
	key, err := strconv.ParseUint(keyStr, 10, 64)
	if err != nil {
		return 0, err
	} else {
		return uint(key), nil
	}
}
